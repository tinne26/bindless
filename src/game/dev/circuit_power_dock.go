package dev

import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

// TODO: update to allow progressive color changes on ephemerous
//       docks. the obvious implementation may slow down some things
//       though.

type PowerDock struct {
	x, y int
	magnet *FloatMagnet
	ephemerousDock uint8
}

func NewPowerDock(col, row int16) *PowerDock {
	x, y := iso.TileCoords(col, row)
	return &PowerDock { x: x, y: y }
}

func (self *PowerDock) PreSetMagnet(magnet *FloatMagnet) {
	self.magnet = magnet
}

func (self *PowerDock) Output() (PolarityType, color.RGBA) {
	if self.magnet == nil { return PolarityNeutral, PolarityNeutral.Color() }
	polarity := self.magnet.Polarity()
	return polarity, polarity.Color()
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
	polarity, clr := self.Output()
	opts.ColorM.ScaleWithColor(clr)
	screen.DrawImage(graphics.DockShape, opts)
	if polarity != PolarityNeutral {
		opts.ColorM.Scale(1, 1, 1, 0.6)
		screen.DrawImage(graphics.DockFill, opts)
	}
}

func (self *PowerDock) MarkEphemerousDock() {
	self.ephemerousDock = 26
}

func (self *PowerDock) Update() {
	if self.ephemerousDock > 0 {
		self.ephemerousDock -= 1
		if self.ephemerousDock == 0 {
			self.magnet = nil
		}
	}
}
