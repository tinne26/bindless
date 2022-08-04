package level

import "image"
import "image/color"
import "fmt"
import "sort"

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/tinne26/etxt"

import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/misc/typewriter"
import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/game/dev"
import "github.com/tinne26/bindless/src/game/sceneitf"
import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/sound"
import "github.com/tinne26/bindless/src/ui"
import "github.com/tinne26/bindless/src/lang"

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
type circuitItf interface {
	Draw(*ebiten.Image)
	Update()
}

type Level struct {
	key levelKey
	fade fadeType
	opacity uint8
	overDir sceneitf.Status 
	tick int
	fadeOutSpeed uint8

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
	circuits iso.Map[circuitItf]
	simUpdateCount int
	abilityExecCues []AbilityExecCue
	pendingRewires []*dev.WireSwitch
	pressingRestart bool

	leftClickPressed bool
	offscreen *ebiten.Image
	renderer *etxt.Renderer
	infoWriter *typewriter.Writer
	horzChoice *ui.HorzChoice
	auxMenu *AuxMenu
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

	lvl := &Level {
		key: key,
		fadeOutSpeed: 3,
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
	}

	// setup writer and choices
	var infoWriter *typewriter.Writer
	text, hasText := levelTexts[int(key)]
	if hasText {
		infoWriter = typewriter.New(text.Get(), ctx)
		infoWriter.SkipToEnd()
		lvl.infoWriter = infoWriter
		lvl.setupHorzChoices(ctx, coda, key)
	}

	// setup aux menu
	rootOpts := []*lang.Text {
		lang.NewText("Restart Level", "Reiniciar Nivel", "Reiniciar Nivell"),
		//lang.NewText("Select Scene", "Seleccionar Escena", "Seleccionar Escena"),
		lang.NewText("Shortcuts", "Atajos", "Dreceres"),
		lang.NewText("Fullscreen", "Resolución", "Resolució"),
		lang.NewText("-- Continue --", "-- Continuar --", "-- Continuar --"),
	}
	auxMenu := ui.NewMenu(ctx, coda, rootOpts, lvl.fnHandlerAuxMenuRoot)
	shortcutsInfo := []*lang.Text {
		lang.NewText("Restart level: TAB", "Reiniciar nivel: TAB", "Reiniciar nivell: TAB"),
		lang.NewText("Skip text: TAB", "Saltar texto: TAB", "Saltar text: TAB"),
		lang.NewText("Select ability: 1-9", "Seleccionar habilidad: 1-9", "Seleccionar habilitat: 1-9"),
		lang.NewText("Fullscreen / windowed: F", "Cambiar resolución: F", "Canviar resolució: F"),
		lang.NewText("Display FPS: D", "Mostrar FPS: D", "Mostrar FPS: D"),
		lang.NewText("Cheat: hold J + scene ID", "Trampas: mantener J + ID escena", "Trampes: aguantar J + ID escena"),
		lang.NewText("-- Back --", "-- Atrás --", "-- Enrere --"),
	}
	auxMenu.SetSubOptions("Shortcuts", shortcutsInfo, lvl.fnHandlerShortcuts)
	auxMenu.SetLogFontSize(12)
	auxMenu.SetLogOptSeparation(22)
	auxMenu.SetLogHorzPadding(2)
	lvl.auxMenu = NewAuxMenu(auxMenu)

	return lvl, nil
}

func (self *Level) Status() sceneitf.Status {
	pressingRestart := misc.SkipKeyPressed()
	if !pressingRestart { self.pressingRestart = false }

	if self.fade == fadeOut {
		if self.opacity == 0 { return self.overDir }
	} else if pressingRestart && !self.pressingRestart {
		sound.PlaySFX(sound.SfxNope)
		return sceneitf.Restart
	}
	return sceneitf.KeepAlive
}

const cycleTicks = 70
func (self *Level) Update(logCursorX, logCursorY int) error {
	// update aux menu
	if self.opacity > 80 {
		wasActive := self.auxMenu.active
		self.auxMenu.Update(logCursorX, logCursorY)
		if self.fade == fadeOut { self.auxMenu.active = false }
		if wasActive || self.auxMenu.active {
			self.pressingRestart = true
			return nil
		}
	}
	
	// update tick
	self.tick += 1
	if self.tick >= cycleTicks { self.tick = 0 }

	// update choices if present
	if self.horzChoice != nil && self.opacity > 80 {
		self.horzChoice.Update(logCursorX, logCursorY)
	}

	// update typewriter if present
	if self.infoWriter != nil {
		self.infoWriter.Tick()
	}

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
		if self.opacity > self.fadeOutSpeed {
			self.opacity -= self.fadeOutSpeed
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
			self.overDir = sceneitf.IsOverNext
			self.fadeOutSpeed = 3
		}
	}

	// update level elements
	self.floatMagnetCount = 0
	if self.tick == 0 {
		// update targets for tutorial level 5
		self.tutorial5Hook()

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
	
	// update circuits (necessary for color transitions)
	self.circuits.Each(
		func(_, _ int16, circuit circuitItf) {
			circuit.Update()
		})
	// update magnets (necessary for color transitions)
	self.magnets.Each(
		func(_, _ int16, magnet dev.Magnet) {
			floatMagnet, isFloatMagnet := magnet.(*dev.FloatMagnet)
			if isFloatMagnet { floatMagnet.Update() }
		})

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
								_ = magnet.Dock(dev.NewTransferProc(transferDock.Source, m.(*dev.FloatMagnet)))
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
	self.circuits.Each(func(_, _ int16, c circuitItf) { c.Draw(self.offscreen) })
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
	self.auxMenu.Draw(self.offscreen)

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
	if self.infoWriter != nil { // tutorial level with information
		fontSize := misc.ScaledFontSize(10, zoomLevel)
		bounds := screen.Bounds()
		xPad, yPad := int(zoomLevel*16), int(zoomLevel*14)
		x, y   := bounds.Min.X + xPad, bounds.Min.Y + yPad
		rect   := image.Rect(x, y, x + int(220*zoomLevel), y + int(320*zoomLevel))
		_, endY := self.infoWriter.Draw(screen, fontSize, rect, self.opacity)
		if self.horzChoice != nil {
			self.horzChoice.SetBaseOpacity(uint8((float64(self.opacity)/255)*128))
			self.horzChoice.SetTopLeft(16, int(float64(endY - bounds.Min.Y)/zoomLevel) + 13)
			self.horzChoice.DrawHiRes(screen, zoomLevel)
		}
	}
	self.auxMenu.DrawHiRes(screen, zoomLevel)
	return // TODO: complete when necessary

	// draw tutorial info and other misc text
	var text string
	if len(self.raisingMagnets) > 0 {
		if self.key == SwitchTutorial {
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
	} else if self.key == CleanerAutomaton {
		text = "(press ESC or TAB to restart the level when you get locked)"
	} else if self.key == SwitchTutorial {
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


// handlers for horzChoice options
func (self *Level) setupHorzChoices(ctx *misc.Context, font *etxt.Font, key levelKey) {
	choiceList := LevelChoices[int(key)]
	if len(choiceList) == 0 { return }
	
	horzChoice := ui.NewHorzChoice(ctx, font)
	horzChoice.SetLogOptSeparation(4)
	for _, choice := range choiceList {
		// link handlers to options. Notice that we are modifying
		// the option objects, but it doesn't matter since they are
		// only used in one scene at a time and without dups
		switch choice.Text.English() {
		case "[ Previous ]":
			if key == Tutorial1 {
				choice.Handler = nil // disabled option
			} else {
				choice.Handler = self.fnPrevHandler
			}
		case "[ Next ]":
			choice.Handler = self.fnNextHandler
		case "[ Solve to continue ]":
			choice.Handler = nil
		case "[ Recharge abilities ]":
			choice.Handler = self.fnRechargeHandler
		default:
			panic(choice.Text.English())
		}
		horzChoice.AddHChoice(choice)
	}
	self.horzChoice = horzChoice
}

func (self *Level) fnNextHandler(_ string) {
	self.fade = fadeOut
	self.fadeOutSpeed = 6
	self.overDir = sceneitf.IsOverNext
	sound.PlaySFX(sound.SfxClick)
	self.horzChoice.Unfocus()
}

func (self *Level) fnPrevHandler(_ string) {
	self.fade = fadeOut
	self.fadeOutSpeed = 6
	self.overDir = sceneitf.IsOverPrev
	sound.PlaySFX(sound.SfxClick)
	self.horzChoice.Unfocus()
}

func (self *Level) fnRechargeHandler(caller string) {
	switch self.key {
	case Tutorial5:
		if caller == "__ongame__" { // magnet reached target normally
			self.abilities.Dock   = 3
			self.abilities.Rewire = 2
			sound.PlaySFX(sound.SfxClick)
		} else { // player manually clicked on the recharge option
			self.abilities.Dock   = 4
			self.abilities.Rewire = 4
			sound.PlaySFX(sound.SfxAbility)
		}
	default:
		panic("unhandled recharge case")
	}
}

// Hacks to make tutorial 5 work with the custom "target" circuit. It moves
// from one position to another when you manage to place a magnet over it,
// and it also partially refills your abilities.
func (self *Level) tutorial5Hook() {
	if self.key != Tutorial5 { return }
	self.circuits.Each(
		func(_, _ int16, circuit circuitItf) {
			devTarget, isTarget := circuit.(*dev.Target)
			if isTarget {
				col, row := devTarget.GetColRow()
				magnet := self.getFloatMagnet(col, row)
				if magnet != nil {
					self.fnRechargeHandler("__ongame__")
					self.circuits.Delete(col, row)
					devTarget.Move()
					col, row := devTarget.GetColRow()
					if col != -1 && row != -1 {
						self.circuits.Set(col, row, devTarget)
					}
				}
			}
		})
}

func (self *Level) fnHandlerAuxMenuRoot(opt string) {
	switch opt {
	case "Restart Level":
		self.fade = fadeOut
		self.opacity = 0
		self.overDir = sceneitf.Restart
		self.auxMenu.menu.Unselect()
		self.auxMenu.active = false
		sound.PlaySFX(sound.SfxClick)
	case "Shortcuts":
		self.auxMenu.menu.NavIn()
		sound.PlaySFX(sound.SfxClick)
	case "Fullscreen":
		self.auxMenu.menu.Unselect()
		self.auxMenu.active = false
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		sound.PlaySFX(sound.SfxClick)
	case "-- Continue --":
		self.auxMenu.active = false
		sound.PlaySFX(sound.SfxClick)
	default:
		panic(opt)
	}
}

func (self *Level) fnHandlerShortcuts(opt string) {
	if opt == "-- Back --" {
		self.auxMenu.menu.NavOut()	
		sound.PlaySFX(sound.SfxClick)
	} else {
		sound.PlaySFX(sound.SfxNope)
	}
}