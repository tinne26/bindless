package dev

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

type PowerDock struct {
	x, y int
	magnet *FloatMagnet
}

func NewPowerDock(col, row int16) *PowerDock {
	x, y := iso.TileCoords(col, row)
	return &PowerDock { x: x, y: y }
}

func (self *PowerDock) PreSetMagnet(magnet *FloatMagnet) {
	self.magnet = magnet
}

func (self *PowerDock) Output() PolarityType {
	if self.magnet == nil { return PolarityNeutral }
	return self.magnet.Polarity()
}

func (self *PowerDock) OnDockChange(magnet *FloatMagnet) {
	if magnet.prevState == StDocking {
		self.magnet = magnet
	} else { // undocking
		self.magnet = nil
	}
}

func (self *PowerDock) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(self.x), float64(self.y))
	screen.DrawImage(graphics.DockShadow, opts)
	polarity := self.Output()
	opts.ColorM.ScaleWithColor(polarity.Color())
	screen.DrawImage(graphics.DockShape, opts)
	if polarity != PolarityNeutral {
		opts.ColorM.Scale(1, 1, 1, 0.6)
		screen.DrawImage(graphics.DockFill, opts)
	}
}
