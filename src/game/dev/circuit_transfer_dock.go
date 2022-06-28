package dev

import "github.com/hajimehoshi/ebiten/v2"

import "bindless/src/game/iso"
import "bindless/src/art/graphics"
import "bindless/src/art/palette"

type TransferDock struct {
	X, Y int
	TargetCol, TargetRow int16
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
	opts.ColorM.ScaleWithColor(palette.PolarityNeutral)
	screen.DrawImage(graphics.DockShape, opts)
}
