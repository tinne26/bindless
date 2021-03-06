package dev

import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"

type StaticMagnet struct {
	polarityFunc func() (PolarityType, color.RGBA)
	x int // tile left x
	y int // tile top y
}

func NewStaticMagnet(col, row int16, polarityFunc func() (PolarityType, color.RGBA)) *StaticMagnet {
	x, y := iso.TileCoords(col, row)
	return &StaticMagnet {
		polarityFunc: polarityFunc,
		x: x,
		y: y,
	}
}

func (self *StaticMagnet) IsAboveHighlight(_ float64) bool { return false }
func (self *StaticMagnet) Update() {} // nothing
func (self *StaticMagnet) LogicalY() int { return self.y }
func (self *StaticMagnet) Polarity() PolarityType {
	polarity, _ := self.polarityFunc()
	return polarity
}
func (self *StaticMagnet) MagneticRange() int16 { return 3 }
func (self *StaticMagnet) Draw(screen *ebiten.Image, _ float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(self.x + 5), float64(self.y + 8))
	screen.DrawImage(graphics.MagnetLargeFloor, opts)
	opts.GeoM.Translate(4, -20)

	polarity, clr := self.polarityFunc()
	opts.ColorM.ScaleWithColor(clr)
	if polarity == PolarityNeutral {
		screen.DrawImage(graphics.MagnetLargeFill, opts)
	} else {
		screen.DrawImage(graphics.MagnetLargeHalo, opts)
	}

	opts.ColorM.Reset()
	screen.DrawImage(graphics.MagnetLarge, opts)

	// TODO: magnetic field animation effects
}
