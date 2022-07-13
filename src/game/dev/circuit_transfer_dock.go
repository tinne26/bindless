package dev

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

type TransferDock struct {
	X, Y int
	TargetCol, TargetRow int16
	ephemerousPower uint8
	polarity PolarityType
}

func NewTransferDockPair(col1, row1 int16, col2, row2 int16) (*TransferDock, *TransferDock) {
	x1, y1 := iso.TileCoords(col1, row1)
	x2, y2 := iso.TileCoords(col2, row2)
	td1 := &TransferDock { X: x1, Y: y1, TargetCol: col2, TargetRow: row2 }
	td2 := &TransferDock { X: x2, Y: y2, TargetCol: col1, TargetRow: row1 }
	return td1, td2
}

func (self *TransferDock) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(self.X), float64(self.Y))
	screen.DrawImage(graphics.DockShadow, opts)
	opts.ColorM.ScaleWithColor(self.polarity.Color())
	screen.DrawImage(graphics.DockShape, opts)
}

func (self *TransferDock) MarkEphemerousPower() {
	self.ephemerousPower = 16
}

func (self *TransferDock) Update() {
	if self.ephemerousPower > 0 {
		self.ephemerousPower -= 1
		if self.ephemerousPower == 0 {
			self.polarity = PolarityNeutral
		}
	}
}

func (self *TransferDock) Output() PolarityType {
	return self.polarity
}
