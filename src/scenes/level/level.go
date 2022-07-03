package level

import "image/color"
import "fmt"
import "sort"

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/tinne26/etxt"

import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/game/dev"
import "github.com/tinne26/bindless/src/game/sceneitf"
import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/sound"

type fadeType uint8
const (
	fadeIn   fadeType = 0
	fadeNone fadeType = 1
	fadeOut  fadeType = 2
)

type ydrawable interface {
	LogicalY() int
	Draw(screen *ebiten.Image, cycle float64) // cycle goes from [0..1) and restarts
}
type drawable interface {
	Draw(*ebiten.Image)
}

type Level struct {
	key levelKey
	fade fadeType
	opacity uint8
	tick int

	baseMapImg *ebiten.Image
	surface iso.Map[struct{}] // only indicates passability
	abilities Abilities
	abilityActiveHint *ebiten.Image
	magnets iso.Map[dev.Magnet]
	highlight *tileHighlight
	fallingMagnets []*dev.FallingMagnet
	raisingMagnets []*dev.RaisingMagnet
	floatMagnetCount int
	winPoints []*dev.WinPoint
	circuits iso.Map[drawable]
	simUpdateCount int
	abilityExecCues []AbilityExecCue
	pendingRewires []*dev.WireSwitch
	pressingRestart bool

	leftClickPressed bool
	offscreen *ebiten.Image
	renderer *etxt.Renderer
}

func New(ctx *misc.Context, key levelKey) (*Level, error) {
	surface := makeLevelSurface(key)
	circuits, magnets := makeLevelDevices(key)
	winPoints := makeLevelWinPoints(key)
	highlight := &tileHighlight{}

	renderer := etxt.NewStdRenderer()
	renderer.SetCacheHandler(ctx.FontCache.NewHandler())
	renderer.SetAlign(etxt.Bottom, etxt.Right)
	renderer.SetQuantizationMode(etxt.QuantizeNone) // TODO: if quantized, bug with cache
	                                                //       not keeping track of quantization
	coda := ctx.FontLib.GetFont("Coda Regular")
	if coda == nil { return nil, fmt.Errorf("missing 'Coda' font") }
	renderer.SetFont(coda)

	return &Level {
		key: key,
		baseMapImg: surface2img(surface),
		surface: surface,
		abilities: makeLevelAbilities(key),
		magnets: magnets,
		highlight: highlight,
		fallingMagnets: make([]*dev.FallingMagnet, 0, 2),
		raisingMagnets: make([]*dev.RaisingMagnet, 0, 1),
		abilityExecCues: make([]AbilityExecCue, 0, 2),
		pendingRewires: make([]*dev.WireSwitch, 0, 1),
		winPoints: winPoints,
		circuits: circuits,
		tick: cycleTicks,
		offscreen: ebiten.NewImage(640, 360),
		renderer: renderer,
		pressingRestart: misc.SkipKeyPressed(),
	}, nil
}

func (self *Level) Status() sceneitf.Status {
	pressingRestart := misc.SkipKeyPressed()
	if !pressingRestart { self.pressingRestart = false }

	if self.fade == fadeOut {
		if self.opacity == 0 { return sceneitf.IsOver }
	} else if pressingRestart && !self.pressingRestart {
		sound.PlaySFX(sound.SfxNope)
		return sceneitf.Restart
	}
	return sceneitf.KeepAlive
}

const cycleTicks = 70
func (self *Level) Update(logCursorX, logCursorY int) error {
	// update tick
	self.tick += 1
	if self.tick >= cycleTicks { self.tick = 0 }

	// update fades
	switch self.fade {
	case fadeIn:
		if self.opacity < 252 {
			self.opacity += 3
		} else {
			self.opacity = 255
			self.fade = fadeNone
		}
	case fadeOut:
		if self.opacity > 3 {
			self.opacity -= 3
		} else {
			self.opacity = 0
		}
	}

	// get tile highlight position
	col, row := iso.CoordsToTile(logCursorX, logCursorY)
	_, found := self.surface.Get(col, row)
	if found {
		self.highlight.active = true
		if self.highlight.col != col || self.highlight.row != row {
			self.highlight.col = col
			self.highlight.row = row
			sound.PlaySFX(sound.SfxLoudNav)
		}
	} else {
		self.highlight.active = false
	}

	// update falling magnets
	i := 0
	for i < len(self.fallingMagnets) {
		canDelete := self.fallingMagnets[i].FallUpdate()
		if canDelete {
			self.fallingMagnets[i] = self.fallingMagnets[len(self.fallingMagnets) - 1]
			self.fallingMagnets = self.fallingMagnets[0 : len(self.fallingMagnets) - 1]
		} else {
			i += 1
		}
	}

	// update raising magnets
	for _, raisingMagnet := range self.raisingMagnets {
		raisingMagnet.Update()
		if raisingMagnet.Y() < 100 && self.fade != fadeOut {
			self.fade = fadeOut
		}
	}

	// update level elements
	self.floatMagnetCount = 0
	if self.tick == 0 {
		// clear ability exec cues
		self.abilityExecCues = self.abilityExecCues[0 : 0]

		// execute pending rewires
		for _, wireSwitch := range self.pendingRewires {
			wireSwitch.Switch()
		}
		self.pendingRewires = self.pendingRewires[0 : 0]

		// detect win conditions
		for _, winPoint := range self.winPoints {
			magnet, found := self.magnets.Get(winPoint.Col, winPoint.Row)
			if found {
				floatMagnet := magnet.(*dev.FloatMagnet)
				pol := floatMagnet.Polarity()
				if pol == winPoint.Polarity {
					self.magnets.Delete(winPoint.Col, winPoint.Row)
					raisingMagnet := floatMagnet.CreateRaising()
					self.raisingMagnets = append(self.raisingMagnets, raisingMagnet)
				}
			}
		}

		// state simulation
		self.magnets.Each(
			func(col, row int16, magnet dev.Magnet) {
				floatMagnet, isFloatMagnet := magnet.(*dev.FloatMagnet)
				if !isFloatMagnet { return }
				self.floatMagnetCount += 1
				fallingMagnet := floatMagnet.StateSim(self.surface, dev.PolarityHack, self.simUpdateCount)
				if fallingMagnet != nil {
					self.magnets.Delete(col, row)
					self.fallingMagnets = append(self.fallingMagnets, fallingMagnet)
				}
			})

		// magnetism simulation
		unsolvedPacks := make([]*dev.CandidateMovesPack, 0)
		self.magnets.Each(
			func(col, row int16, magnet dev.Magnet) {
				floatMagnet, isFloatMagnet := magnet.(*dev.FloatMagnet)
				if !isFloatMagnet { return }
				candidateMoves := floatMagnet.MagneticSim(self.magnets)
				if candidateMoves != nil {
					if candidateMoves.Empty() { panic("empty candidate moves!") }
					unsolvedPacks = append(unsolvedPacks, candidateMoves)
				}
			})

		// the heavy part
		moveChoices := dev.SolveCandidateMovesSystem(unsolvedPacks, self.magnets)
		for _, moveChoice := range moveChoices {
			magnet := moveChoice.Magnet
			cmpMagnet, found := self.magnets.Get(magnet.Column(), magnet.Row())
			if found && cmpMagnet == magnet {
				self.magnets.Delete(magnet.Column(), magnet.Row())
			}
			self.magnets.Set(moveChoice.TargetColumn, moveChoice.TargetRow, magnet)
			magnet.ConfirmMove(moveChoice.TargetColumn, moveChoice.TargetRow)
		}

		self.simUpdateCount += 1
	} else { // still count floating magnets
		self.magnets.Each(
			func(_, _ int16, magnet dev.Magnet) {
				_, isFloatMagnet := magnet.(*dev.FloatMagnet)
				if isFloatMagnet { self.floatMagnetCount += 1 }
			})
	}

	// detect left-clicks for interaction
	leftClickPressed := misc.MousePressed()
	newClick := false
	if leftClickPressed && !self.leftClickPressed {
		self.leftClickPressed = true
		newClick = true
	} else {
		self.leftClickPressed = leftClickPressed
	}

	// update abilities (clicking on abilities)
	self.abilities.Update(newClick, logCursorX, logCursorY)
	if newClick && self.highlight.active && self.abilities.Selected == 0 {
		self.abilities.BlinkLeft = 6
		sound.PlaySFX(sound.SfxNope)
	}

	// handle hovering state and apply clicks if relevant
	self.abilityActiveHint = nil
	if self.highlight.active && self.abilities.Selected != 0 {
		col, row := self.highlight.col, self.highlight.row
		abilityUsed := false
		switch self.abilities.Selected {
		case 1: // dock
			circuit, hasCircuit := self.circuits.Get(col, row)
			if hasCircuit {
				magnet := self.getFloatMagnet(col, row)
				transferDock, hasTransferDock := circuit.(*dev.TransferDock)
				powerDock, hasPowerDock := circuit.(*dev.PowerDock)
				if magnet != nil && (hasTransferDock || hasPowerDock) {
					if magnet.CanDock() {
						self.abilityActiveHint = graphics.HudDock
						if newClick {
							abilityUsed = true
							if hasTransferDock {
								m, f := self.magnets.Get(transferDock.TargetCol, transferDock.TargetRow)
								if !f { panic("no magnet placed at transfer dock!") }
								_ = magnet.Dock(m.(*dev.FloatMagnet))
							} else { // hasPowerDock
								_ = magnet.Dock(powerDock)
							}
						}
					} else if magnet.CanUndock() {
						abilityUsed = true
						self.abilityActiveHint = graphics.HudUndock
						if newClick { magnet.Undock() }
					}
				}
			}
		case 2: // rewire
			circuit, hasCircuit := self.circuits.Get(col, row)
			if hasCircuit {
				wireSwitch, isSwitch := circuit.(*dev.WireSwitch)
				if isSwitch && wireSwitch.CanSwitch() {
					self.abilityActiveHint = graphics.HudRewire
					if newClick {
						abilityUsed = true
						wireSwitch.SetPendingSwitch()
						self.pendingRewires = append(self.pendingRewires, wireSwitch)
					}
				}
			}
		case 3: // switch
			magnet := self.getFloatMagnet(col, row)
			if magnet != nil && magnet.CanSwitch() {
				self.abilityActiveHint = graphics.HudSwitch
				if newClick {
					abilityUsed = true
					magnet.Switch()
				}
			}
		}

		if newClick {
			if abilityUsed {
				self.abilities.ConsumeCharge()
				self.abilityExecCues = append(self.abilityExecCues, NewAbilityExecCue(col, row))
				sound.PlaySFX(sound.SfxAbility)
			} else {
				self.abilities.BlinkLeft = 6
				sound.PlaySFX(sound.SfxNope)
			}
		}
	}

	// update highlight cutting state, as previous updates can lead to changes
	if self.highlight.active {
		self.highlight.cutting = false
		magnet, found := self.magnets.Get(col, row)
		if found {
			self.highlight.cutting = !magnet.IsAboveHighlight(float64(self.tick)/cycleTicks)
		}
	}

	return nil
}

func (self *Level) Draw(screen *ebiten.Image) {
	cycle := float64(self.tick)/cycleTicks
	self.offscreen.Clear()

	// draw tile map
	self.offscreen.DrawImage(self.baseMapImg, nil)

	// draw circuits
	self.circuits.Each(func(_, _ int16, c drawable) { c.Draw(self.offscreen) })
	for _, winPoint := range self.winPoints { winPoint.Draw(self.offscreen) }

	// TODO: draw floating tiles

	// draw magnets and tile highlight
	sortedElems := make([]ydrawable, 0, self.magnets.Size() + 1)
	self.magnets.Each(func(_, _ int16, magnet dev.Magnet) {
		sortedElems = append(sortedElems, magnet)
	})
	if self.highlight.active {
		sortedElems = append(sortedElems, self.highlight)
	}
	for _, fallingMagnet := range self.fallingMagnets {
		sortedElems = append(sortedElems, fallingMagnet)
	}

	sort.Slice(sortedElems, func(i, j int) bool {
		return sortedElems[i].LogicalY() < sortedElems[j].LogicalY()
	})
	for _, elem := range sortedElems {
		elem.Draw(self.offscreen, cycle)
	}

	// draw raising magnets, which go above everything else (due to
	// level design, not by actual generalizable behavior)
	for _, raisingMagnet := range self.raisingMagnets {
		raisingMagnet.Draw(self.offscreen)
	}

	// draw HUD
	self.abilities.Draw(self.offscreen)

	for _, execCue := range self.abilityExecCues {
		execCue.Draw(self.offscreen, cycle)
	}

	// draw visual hints for abilities
	if self.abilityActiveHint != nil {
		self.abilities.DrawWordHint(self.offscreen, self.highlight.col, self.highlight.row, self.abilityActiveHint)
	}

	// draw offscreen to main screen
	if self.opacity == 255 {
		screen.DrawImage(self.offscreen, nil)
	} else {
		opts := &ebiten.DrawImageOptions{}
		opts.ColorM.Scale(1.0, 1.0, 1.0, float64(self.opacity)/255)
		screen.DrawImage(self.offscreen, opts)
	}
}

func (self *Level) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	// draw tutorial info and other misc text
	var text string
	if len(self.raisingMagnets) > 0 {
		if self.key == SwitchTest {
			text = "Tutorial level complete..."
		} else {
			text = "Disabling MSP security layer..."
		}
	} else if self.floatMagnetCount == 0 {
		text = "(press ESC or TAB to restart the level)"
	} else if self.key == CleanerTestDock {
		if self.abilities.Dock > 0 {
			if self.abilities.Selected != 1 {
				text = "(click the first icon on the bottom left to select the 'Dock' ability)"
			} else {
				text = "(click on the tile below the floating magnet to use the selected ability)"
			}
		}
	} else if self.key == CleanerTestRewire {
		if self.abilities.Rewire > 0 {
			if self.abilities.Selected != 2 {
				text = "(select the 'Rewire' ability. Numeric keys also work (try 2))"
			} else {
				text = "(use 'Rewire' on the tile where the wires split)"
			}
		}
	} else if self.key == CleanerTestReal {
		text = "(press ESC or TAB to restart the level when you get locked)"
	} else if self.key == SwitchTest {
		if self.abilities.Switch > 0 {
			if self.abilities.Selected != 3 {
				text = "('Switch' allows changing the polarity of small magnets)"
			} else {
				text = "('Switch' can be used on small magnets even if they are docked)"
			}
		} else {
			magnet := self.getFloatMagnet(16, 20)
			if magnet != nil {
				if magnet.Polarity() != dev.PolarityNegative && !magnet.HasPendingSwitch() {
					text = "(yeah, that won't work, but kudos for experimenting ^^)"
				}
			}
		}
	}

	if text != "" {
		if self.opacity <= 95 { return }
		self.renderer.SetColor(color.RGBA{255, 255, 255, self.opacity - 95})
		self.renderer.SetTarget(screen)
		self.renderer.SetSizePx(misc.ScaledFontSize(10, zoomLevel))

		bounds := screen.Bounds()
		self.renderer.Draw(text, bounds.Max.X - int(10*zoomLevel), bounds.Max.Y - int(10*zoomLevel))
	}
}

func (self *Level) getFloatMagnet(col, row int16) *dev.FloatMagnet {
	magnet, found := self.magnets.Get(col, row)
	if !found { return nil }
	floatMagnet, isFloatMagnet := magnet.(*dev.FloatMagnet)
	if !isFloatMagnet { return nil }
	return floatMagnet
}
