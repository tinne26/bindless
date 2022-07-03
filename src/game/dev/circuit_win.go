package dev

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

type WinPoint struct {
	Col, Row int16
	Polarity PolarityType
}

func (self *WinPoint) Draw(screen *ebiten.Image) {
	x, y := iso.TileCoords(self.Col, self.Row)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x + 0), float64(y + 0))
	screen.DrawImage(graphics.FieldShadow, opts)
	opts.ColorM.ScaleWithColor(self.Polarity.Color())
	screen.DrawImage(graphics.FieldShape, opts)
}
