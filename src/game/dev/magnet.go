package dev

import "image"
import "math/rand"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

type Magnet interface {
	Polarized
	MagneticRange() int16
	Draw(screen *ebiten.Image, cycle float64)
	LogicalY() int
	IsAboveHighlight(cycle float64) bool
}

type dockChangeHandler interface {
	OnDockChange(*FloatMagnet)
}

type FloatMagnetState uint8
const (
	StDocked    FloatMagnetState = 0
	StDocking   FloatMagnetState = 1
	StUndocking FloatMagnetState = 2
	StFloating  FloatMagnetState = 3
	StMoveNE FloatMagnetState = 4
	StMoveNW FloatMagnetState = 5
	StMoveSE FloatMagnetState = 6
	StMoveSW FloatMagnetState = 7
)

type FloatMagnet struct {
	polarity PolarityType
	dockChangeHandler dockChangeHandler
	prevState FloatMagnetState
	state FloatMagnetState
	nextState FloatMagnetState
	spectreSimsLeft int
	col int16 // tile col
	row int16 // tile row
	pendingSpectre bool
	pendingPolaritySwitch bool
	sway float64 // between 0.0 and 0.2
	nextSwayTarget float64
	currentSimUpdate int
}

func NewFloatMagnet(col, row int16, state FloatMagnetState, polarity PolarityType) *FloatMagnet {
	if state != StDocked && state != StFloating {
		panic("initial float magnet state can only be docked or floating")
	}
	sway := rand.Float64()/5
	return &FloatMagnet {
		prevState: state,
		state: state,
		nextState: state,
		polarity: polarity,
		col: col,
		row: row,
		sway: sway,
	}
}

func (self *FloatMagnet) Column() int16 { return self.col }
func (self *FloatMagnet) Row()    int16 { return self.row }

func (self *FloatMagnet) LogicalY() int { return iso.YCoord(self.col, self.row) }
func (self *FloatMagnet) MagneticRange() int16 { return 2 }
func (self *FloatMagnet) Polarity() PolarityType {
	if self.spectreSimsLeft > 0 { return PolarityNeutral }
	return self.polarity
}

func (self *FloatMagnet) HasPendingSwitch() bool {
	return self.pendingPolaritySwitch
}

func (self *FloatMagnet) CanUndock() bool {
	if self.state != StDocked && self.nextState != StDocked { return false }
	if self.polarity == PolarityNeutral { return false }
	if self.dockChangeHandler == nil { return true }
	_, inDockTransfer := self.dockChangeHandler.(*FloatMagnet)
	return !inDockTransfer
}
func (self *FloatMagnet) Undock() bool {
	if !self.CanUndock() { return false }
	self.nextState = StUndocking
	return true
}
func (self *FloatMagnet) CanDock() bool {
	return self.nextState == StFloating
}
func (self *FloatMagnet) Dock(dch dockChangeHandler) bool {
	if !self.CanDock() { return false }
	self.nextState = StDocking
	self.dockChangeHandler = dch
	return true
}

func (self *FloatMagnet) PreSetDockChangeHandler(dch dockChangeHandler) {
	self.dockChangeHandler = dch
}

func (self *FloatMagnet) OnDockChange(magnet *FloatMagnet) {
	// magnet just docked, and self has to undock
	if self.currentSimUpdate >= magnet.currentSimUpdate {
		// self already updated state in this round, so we have
		// to fully correct self state
		self.polarity = magnet.Polarity()
		self.prevState = StDocked
		self.state = StUndocking
		self.nextState = StFloating
	} else {
		// self hasn't updated state in this round, so we
		// configure self state so it will be updated naturally
		self.polarity = magnet.Polarity()
		self.prevState = self.state
		self.state = StDocked
		self.nextState = StUndocking
	}
}

func (self *FloatMagnet) CanSwitch() bool {
	return self.polarity != PolarityNeutral && !self.pendingPolaritySwitch && !self.pendingSpectre
}
func (self *FloatMagnet) Switch() bool {
	if !self.CanSwitch() { return false }
	self.pendingPolaritySwitch = true
	return true
}
func (self *FloatMagnet) CanSpectre() bool {
	return !self.pendingPolaritySwitch && !self.pendingSpectre
}
func (self *FloatMagnet) Spectre() bool {
	if !self.CanSpectre() { return false }
	self.pendingSpectre = true
	return true
}

func (self *FloatMagnet) StateSim(surface iso.Map[struct{}], floatTilePolarity PolarityType, simUpdateCount int) *FallingMagnet {
	// hack for tricky situations
	self.currentSimUpdate = simUpdateCount

	// adjust sway so magnets don't all hover still in perfect sync, which is ugly
	if self.nextSwayTarget == 0 {
		self.nextSwayTarget = rand.Float64()/5
	}
	if self.nextSwayTarget > self.sway {
		self.sway += 0.04
		if self.sway > self.nextSwayTarget {
			self.sway = self.nextSwayTarget
			self.nextSwayTarget = 0
		}
	} else {
		self.sway -= 0.04
		if self.sway < self.nextSwayTarget {
			self.sway = self.nextSwayTarget
			self.nextSwayTarget = 0
		}
	}

	// reduce spectre-mode turns left
	if self.spectreSimsLeft > 0 {
		self.spectreSimsLeft -= 1
	}

	// switch polarity if required
	if self.pendingPolaritySwitch {
		self.pendingPolaritySwitch = false
		if self.polarity == PolarityPositive {
			self.polarity = PolarityNegative
		} else {
			self.polarity = PolarityPositive
		}
	} else if self.pendingSpectre {
		self.pendingSpectre = false
		self.spectreSimsLeft = 6
	}

	// determine if this magnet must fall through the current tile
	// (to not fall, there must be a surface under us or the floating
   //  tile polarity must match ours)
	// (current tile is floating tile and spectre or opposite polarity)
	_, overSurface := surface.Get(self.col, self.row)
	if !overSurface {
		if self.spectreSimsLeft > 0 || floatTilePolarity != self.polarity {
			return self.createFalling()
		}
	}

	// natural state transition
	self.prevState = self.state
	self.state = self.nextState
	switch self.state {
	case StDocked:
		self.nextState = StDocked
		if self.prevState == StDocking {
			self.dockChangeHandler.OnDockChange(self)
			_, isFloatMagnet := self.dockChangeHandler.(*FloatMagnet)
			if isFloatMagnet {
				self.polarity = PolarityNeutral
				self.dockChangeHandler = nil
				// TODO: ^ I moved this line inside the if recently
				//         in case something breaks
			}
		}
	case StDocking:
		self.nextState = StDocked
	case StUndocking:
		self.nextState = StFloating
		if self.dockChangeHandler != nil { // notify power dock case
			powDock := self.dockChangeHandler.(*PowerDock)
			powDock.OnDockChange(self)
			if self.prevState == StDocking { // edge case of consecutive dock/undock
				powDock.MarkEphemerousDock()
			}
			self.dockChangeHandler = nil
		}
	case StFloating:
		self.nextState = StFloating
	default:
		panic(self.state)
	}

	return nil
}

// this is the hell we never wanted but deserved
func (self *FloatMagnet) MagneticSim(magnets iso.Map[Magnet]) *CandidateMovesPack {
	// handle spectre mode separately
	if self.spectreSimsLeft > 0 && self.prevState != StUndocking && self.prevState != StFloating {
		// see if we can continue (only check collisions, anything else is
		// fair game, let the spectre power do weird stuff and throw magnets
		// over and outside the universe)
		if self.couldMoveDir(magnets, self.prevState) { // may continue in current direction
			candidateMoves := &CandidateMovesPack{ Magnet: self }
			switch self.prevState {
			case StMoveNE: candidateMoves.NE = 1
			case StMoveNW: candidateMoves.NW = 1
			case StMoveSE: candidateMoves.SE = 1
			case StMoveSW: candidateMoves.SW = 1
			}
			return candidateMoves
		} else { // stay still in spectre
			return nil
		}
	}

	// abort if not floating, we can't trigger moves
	if self.state < StFloating { return nil }

	// we are floating and are not in spectre, apply magnetism for movement
	const ScanDist = 3
	candidateMoves := &CandidateMovesPack{ Magnet: self }
	canMoveNE := self.couldMoveDir(magnets, StMoveNE)
	canMoveSW := self.couldMoveDir(magnets, StMoveSW)
	canMoveNW := self.couldMoveDir(magnets, StMoveNW)
	canMoveSE := self.couldMoveDir(magnets, StMoveSE)
	onLock := false

	// NE scan
	for i := int16(1); i <= ScanDist; i++ {
		magnet, found := magnets.Get(self.col + i, self.row)
		if !found || magnet.Polarity() == PolarityNeutral { continue }
		if i > magnet.MagneticRange() { continue }
		if magnet.Polarity() != self.polarity {
			if i == 1 { onLock = true }
			if (candidateMoves.NE == 0 || i < candidateMoves.NE) && canMoveNE { candidateMoves.NE = i }
		} else {
			if (candidateMoves.SW == 0 || i < candidateMoves.SW) && canMoveSW { candidateMoves.SW = i }
		}
	}

	// NW scan
	for i := int16(1); i <= ScanDist; i++ {
		magnet, found := magnets.Get(self.col, self.row - i)
		if !found || magnet.Polarity() == PolarityNeutral { continue }
		if i > magnet.MagneticRange() { continue }
		if magnet.Polarity() != self.polarity {
			if i == 1 { onLock = true }
			if (candidateMoves.NW == 0 || i < candidateMoves.NW) && canMoveNW { candidateMoves.NW = i }
		} else {
			if (candidateMoves.SE == 0 || i < candidateMoves.SE) && canMoveSE { candidateMoves.SE = i }
		}
	}

	// SE scan
	for i := int16(1); i <= ScanDist; i++ {
		magnet, found := magnets.Get(self.col, self.row + i)
		if !found || magnet.Polarity() == PolarityNeutral { continue }
		if i > magnet.MagneticRange() { continue }
		if magnet.Polarity() != self.polarity {
			if i == 1 { onLock = true }
			if (candidateMoves.SE == 0 || i < candidateMoves.SE) && canMoveSE { candidateMoves.SE = i }
		} else {
			if (candidateMoves.NW == 0 || i < candidateMoves.NW) && canMoveNW { candidateMoves.NW = i }
		}
	}

	// SW scan
	for i := int16(1); i <= ScanDist; i++ {
		magnet, found := magnets.Get(self.col - i, self.row)
		if !found || magnet.Polarity() == PolarityNeutral { continue }
		if i > magnet.MagneticRange() { continue }
		if magnet.Polarity() != self.polarity {
			if i == 1 { onLock = true }
			if (candidateMoves.SW == 0 || i < candidateMoves.SW) && canMoveSW { candidateMoves.SW = i }
		} else {
			if (candidateMoves.NE == 0 || i < candidateMoves.NE) && canMoveNE { candidateMoves.NE = i }
		}
	}

	if candidateMoves.Empty() { return nil }
	candidateMoves.CorrectOpposing()
	if candidateMoves.Empty() { return nil }
	if onLock && !candidateMoves.CanUnlock() { return nil }
	candidateMoves.ApplyInertia(self.prevState)
	return candidateMoves
}

func minInt16(a, b int16) int16 { if a >= b { return a } else { return b } }

func (self *FloatMagnet) ConfirmMove(col, row int16) {
	validMove := false
	if col == self.col {
		if row == self.row + 1 {
			self.state = StMoveSE
			validMove  = true
		} else if row == self.row - 1 {
			self.state = StMoveNW
			validMove  = true
		}
	} else if row == self.row {
		if col == self.col + 1 {
			self.state = StMoveNE
			validMove  = true
		} else if col == self.col - 1 {
			self.state = StMoveSW
			validMove  = true
		}
	}

	if !validMove { panic("terrible assumptions, tinne") }
	self.col, self.row = col, row
}

func (self *FloatMagnet) IsAboveHighlight(cycle float64) bool {
	switch self.state {
	case StDocked:
		return false
	case StDocking:
		return cycle < 0.8
	case StUndocking:
		return cycle > 0.2
	default:
		return true
	}
}

func (self *FloatMagnet) Draw(screen *ebiten.Image, cycle float64) {
	x, y, shadowY, bframe := self.currentDrawParams(cycle)

	// draw shadow if necessary
	if self.state != StDocked {
		offset := 0
		if !bframe { offset = 6 }
		aniRect := image.Rect(0 + offset, 0, 6 + offset, 3)
		shadow := graphics.MagnetSmallShadowAni.SubImage(aniRect).(*ebiten.Image)
		opts := &ebiten.DrawImageOptions{}
		opts.CompositeMode = ebiten.CompositeModeSourceAtop
		opts.GeoM.Translate(float64(x + 5), float64(shadowY))
		screen.DrawImage(shadow, opts)
		opts.CompositeMode = ebiten.CompositeModeSourceOver
	}

	if bframe && self.state >= StFloating { y += 1 }
	drawSmallMagnetAt(screen, x, y, self.polarity, self.spectreSimsLeft > 0)
}

func (self *FloatMagnet) couldMoveDir(magnets iso.Map[Magnet], move FloatMagnetState) bool {
	newCol, newRow := applyMoveToPos(self.col, self.row, move)
	obstacle, found := magnets.Get(newCol, newRow)
	if !found { return true }
	switch mgnt := obstacle.(type) {
	case *StaticMagnet:
		return false
	case *FloatMagnet:
		if mgnt.state == StDocked || mgnt.state == StDocking { return false }
		return true
	default:
		panic("nope")
	}
}

func (self *FloatMagnet) currentDrawParams(cycle float64) (int, int, int, bool) {
	const xMoveChange = 18
	const yMoveChange = 9

	bframe := (cycle >= 0.3 - self.sway && cycle <= 0.7 - self.sway)
	x, y := iso.TileCoords(self.col, self.row)
	x += 9
	y -= 11
	shadowFix := 0

	switch self.state {
	case StDocked: // docked
		y += 8
	case StDocking: // docking
		change := int(8*cycle)
		y += change
		shadowFix = -change
	case StUndocking: // undocking
		change := 8 - int(8*cycle)
		y += change
		shadowFix = -change
	case StFloating: // floating
		// y change is applied later in draws to not collide with shadow
		// position. can it get hackier? yeeees, do not challenge me!
	case StMoveNE:
		x -= xMoveChange - int(xMoveChange*cycle)
		y += yMoveChange - int(yMoveChange*cycle)
	case StMoveNW:
		x += xMoveChange - int(xMoveChange*cycle)
		y += yMoveChange - int(yMoveChange*cycle)
	case StMoveSE:
		x -= xMoveChange - int(xMoveChange*cycle)
		y -= yMoveChange - int(yMoveChange*cycle)
	case StMoveSW:
		x += xMoveChange - int(xMoveChange*cycle)
		y -= yMoveChange - int(yMoveChange*cycle)
	default: // moving
		panic("invalid floating magnet state")
	}

	return x, y, y + 18 + shadowFix, bframe
}

func (self *FloatMagnet) createFalling() *FallingMagnet {
	x, y, _, _ := self.currentDrawParams(1.0)
	return &FallingMagnet {
		inSpectre: self.spectreSimsLeft > 0,
		polarity: self.polarity,
		y: float64(y),
		origX: x,
		origY: y,
		speed: 0.2,
	}
}

func (self *FloatMagnet) CreateRaising() *RaisingMagnet {
	x, y, _, _ := self.currentDrawParams(1.0)
	return &RaisingMagnet{ polarity: self.polarity, x: x, y: y }
}

func drawSmallMagnetAt(screen *ebiten.Image, x int, y int, polarity PolarityType, inSpectre bool) {
	// TODO: handle inSpectre case
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	opts.ColorM.ScaleWithColor(polarity.Color())
	if polarity == PolarityNeutral {
		screen.DrawImage(graphics.MagnetSmallFill, opts)
	} else {
		screen.DrawImage(graphics.MagnetSmallHalo, opts)
	}
	screen.DrawImage(graphics.MagnetSmall, opts)
}

func applyMoveToPos(col, row int16, move FloatMagnetState) (int16, int16) {
	switch move {
	case StMoveNE: return col + 1, row
	case StMoveNW: return col, row - 1
	case StMoveSE: return col, row + 1
	case StMoveSW: return col - 1, row
	default:
		panic("applyMoveToPos received non-move state")
	}
}
