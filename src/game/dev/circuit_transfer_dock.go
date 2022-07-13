package dev

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

type TransferDock struct {
	X, Y int
	TargetCol, TargetRow int16
	Source *TransferSource
}

func NewTransferDockPair(col1, row1 int16, col2, row2 int16) (*TransferDock, *TransferDock) {
	x1, y1 := iso.TileCoords(col1, row1)
	x2, y2 := iso.TileCoords(col2, row2)
	td1 := &TransferDock { X: x1, Y: y1, TargetCol: col2, TargetRow: row2 }
	td2 := &TransferDock { X: x2, Y: y2, TargetCol: col1, TargetRow: row1 }
	source := newTransferSource(td1, td2)
	td1.Source = source
	td2.Source = source
	return td1, td2
}

func (self *TransferDock) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(self.X), float64(self.Y))
	screen.DrawImage(graphics.DockShadow, opts)
	opts.ColorM.ScaleWithColor(self.Source.color)
	screen.DrawImage(graphics.DockShape, opts)
}

func (self *TransferDock) Update() {
	self.Source.Update()
}
