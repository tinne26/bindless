package dev

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"


type WireConn uint8
const (
	ConnNW WireConn = 0
	ConnNE WireConn = 1
	ConnSW WireConn = 2
	ConnSE WireConn = 3
)

type Wire2 struct { x, y int; a, b WireConn; polaritySrcFunc func() PolarityType }
func NewWire2(col, row int16, a, b WireConn, polaritySrcFunc func() PolarityType) *Wire2 {
	if a == b { panic("invalid wire connection") }
	x, y := iso.TileCoords(col, row)
	return &Wire2 { x, y, a, b, polaritySrcFunc }
}

func (self *Wire2) Update() {} // nothing for this type of circuit

func (self *Wire2) Draw(screen *ebiten.Image) {
	drawWire2(screen, self.x, self.y, self.polaritySrcFunc(), self.a, self.b)
}

func drawWire2(screen *ebiten.Image, x, y int, polarity PolarityType, a, b WireConn) {
	if b < a { a, b = b, a }

	var src [2]*ebiten.Image
	switch a {
	case ConnNW:
		switch b {
		case ConnNE: src = graphics.WireNW2NE
		case ConnSW: src = graphics.WireNW2SW
		case ConnSE: src = graphics.WireNW2SE
		}
	case ConnNE:
		switch b {
		case ConnSW: src = graphics.WireNE2SW
		case ConnSE: src = graphics.WireNE2SE
		}
	case ConnSW:
		src = graphics.WireSW2SE
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x + 8), float64(y + 4))
	screen.DrawImage(src[1], opts) // draw shadow first
	opts.ColorM.ScaleWithColor(polarity.Color())
	screen.DrawImage(src[0], opts) // draw wire with proper color
}
