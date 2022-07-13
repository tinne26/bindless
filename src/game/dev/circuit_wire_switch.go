package dev

import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/palette"

type WireSwitch struct {
	x, y int
	src WireConn
	a, b WireConn
	switched bool // if false, we go to a, if true, we go to b
	pendingSwitch bool
	polaritySrcFunc func() PolarityType
}

func NewWireSwitch(col, row int16, src, a, b WireConn, polaritySrcFunc func() PolarityType) *WireSwitch {
	if a == b || src == a || src == b { panic("invalid wire connection") }
	x, y := iso.TileCoords(col, row)
	return &WireSwitch { x: x, y: y, src: src, a: a, b: b, switched: false, polaritySrcFunc: polaritySrcFunc }
}

func (self *WireSwitch) Update() {} // nothing for this type of circuit

func (self *WireSwitch) OutA() PolarityType {
	if self.switched { return PolarityNeutral }
	return self.polaritySrcFunc()
}

func (self *WireSwitch) OutB() PolarityType {
	if !self.switched { return PolarityNeutral }
	return self.polaritySrcFunc()
}

func (self *WireSwitch) SetPendingSwitch() {
	if self.pendingSwitch { panic("attempted to set pending switch when one already exists") }
	self.pendingSwitch = true
}

func (self *WireSwitch) Switch() {
	if !self.pendingSwitch { panic("attempted to switch without pending switch") }
	self.pendingSwitch = false
	self.switched = !self.switched
}

func (self *WireSwitch) CanSwitch() bool {
	return !self.pendingSwitch
}

func (self *WireSwitch) Draw(screen *ebiten.Image) {
	if !self.switched {
		drawWire2(screen, self.x, self.y, self.polaritySrcFunc(), self.src, self.a)
		drawWireEnd(screen, self.x, self.y, self.b)
	} else {
		drawWire2(screen, self.x, self.y, self.polaritySrcFunc(), self.src, self.b)
		drawWireEnd(screen, self.x, self.y, self.a)
	}
}

var pxImg *ebiten.Image
var pxImgNeutral *ebiten.Image
func drawWireEnd(screen *ebiten.Image, x, y int, conn WireConn) {
	if pxImg == nil {
		pxImg = ebiten.NewImage(1, 1)
		pxImg.Fill(color.White)
	}
	opts := &ebiten.DrawImageOptions{}

	x += 8 // move from tile coord to wire rect coord
	y += 4

	// don't ask please
	switch conn {
	case ConnNE:
		x += 14
		opts.GeoM.Translate(float64(x), float64(y + 1))
		opts.ColorM.ScaleWithColor(palette.PolarityNeutral)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(2, -1)
		screen.DrawImage(pxImg, opts)

		opts.GeoM.Translate(-2, 0)
		opts.ColorM.ScaleWithColor(palette.CircuitShadow)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(1, 1)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(2, -1)
		screen.DrawImage(pxImg, opts)
	case ConnNW:
		x += 1
		opts.GeoM.Translate(float64(x), float64(y))
		opts.ColorM.ScaleWithColor(palette.PolarityNeutral)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(2, 1)
		screen.DrawImage(pxImg, opts)

		opts.GeoM.Translate(-3, -1)
		opts.ColorM.ScaleWithColor(palette.CircuitShadow)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(2, 1)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(1, -1)
		screen.DrawImage(pxImg, opts)
	case ConnSE:
		x += 15
		y += 7
		opts.GeoM.Translate(float64(x), float64(y))
		opts.ColorM.ScaleWithColor(palette.PolarityNeutral)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(2, 1)
		screen.DrawImage(pxImg, opts)

		opts.GeoM.Translate(-3, -1)
		opts.ColorM.ScaleWithColor(palette.CircuitShadow)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(1, -1)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(1,  2)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(1, -1)
		screen.DrawImage(pxImg, opts)
	case ConnSW:
		y += 7
		opts.GeoM.Translate(float64(x), float64(y + 1))
		opts.ColorM.ScaleWithColor(palette.PolarityNeutral)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(2, -1)
		screen.DrawImage(pxImg, opts)

		opts.GeoM.Translate(-1, 1)
		opts.ColorM.ScaleWithColor(palette.CircuitShadow)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(-1, -1)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(3,  0)
		screen.DrawImage(pxImg, opts)
		opts.GeoM.Translate(-1, -1)
		screen.DrawImage(pxImg, opts)
	}
}
