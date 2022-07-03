package graphics

import "image"
import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/art/palette"

// couldn't you just create some nice images? no...
func loadWires() {
	white := color.RGBA{255, 255, 255, 255} // for masking
	shadw := palette.CircuitShadow

	// simple wires
	WireNW2NE[0] = makeWireImg(white, 1, 0,  3, 1,       5, 2,       7, 3,             10, 3,              12, 2,        14, 1,        16, 0)
	WireNW2NE[1] = makeWireImg(shadw, 0, 0,  2, 1, 3, 0, 4, 2, 5, 1, 6, 3, 7, 2, 8, 3, 11, 3, 10, 2, 9, 3, 13, 2, 12, 1, 15, 1, 14, 0, 17, 0)

	WireNW2SW[0] = makeWireImg(white, 1, 0,  3, 1,       5, 2,       7, 3,                        6, 5,       4, 6,       2, 7,       0, 8)
	WireNW2SW[1] = makeWireImg(shadw, 0, 0,  2, 1, 3, 0, 4, 2, 5, 1, 6, 3, 7, 2, 8, 3,      7, 4, 6, 4, 7, 5, 4, 5, 5, 6, 2, 6, 3, 7, 0, 7, 1, 8)

	WireNW2SE[0] = makeWireImg(white, 1, 0,  3, 1,       5, 2,       7, 3,       9, 4,       11, 5,        13, 6,        15, 7,        17, 8)
	WireNW2SE[1] = makeWireImg(shadw, 0, 0,  2, 1, 3, 0, 4, 2, 5, 1, 6, 3, 7, 2, 8, 4, 9, 3, 10, 5, 11, 4, 12, 6, 13, 5, 14, 7, 15, 6, 16, 8, 17, 7)

	WireNE2SW[0] = makeWireImg(white, 16, 0,  14, 1,        12, 2,        10, 3,            8, 4,             6, 5,       4, 6,       2, 7,       0, 8)
	WireNE2SW[1] = makeWireImg(shadw, 17, 0,  15, 1, 14, 0, 13, 2, 12, 1, 11, 3, 10, 2,     9, 4, 8, 3,       6, 4, 7, 5, 4, 5, 5, 6, 2, 6, 3, 7, 0, 7, 1, 8)

	WireNE2SE[0] = makeWireImg(white, 16, 0,  14, 1,        12, 2,        10, 3,                           11, 5,        13, 6,        15, 7,        17, 8)
	WireNE2SE[1] = makeWireImg(shadw, 17, 0,  15, 1, 14, 0, 13, 2, 12, 1, 11, 3, 10, 2,    9, 3, 10, 4,    10, 5, 11, 4, 12, 6, 13, 5, 14, 7, 15, 6, 16, 8, 17, 7)

	WireSW2SE[0] = makeWireImg(white, 0, 8,       2, 7,       4, 6,       6, 5,          8, 4, 9, 4,       11, 5,        13, 6,        15, 7,        17, 8)
	WireSW2SE[1] = makeWireImg(shadw, 0, 7, 1, 8, 2, 6, 3, 7, 4, 5, 5, 6, 6, 4, 7, 5,    8, 3, 9, 3,       11, 4, 12, 6, 13, 5, 14, 7, 15, 6, 16, 8, 17, 7)
}

func makeWireImg(clr color.RGBA, pxCoords ...int) *ebiten.Image {
	img := image.NewRGBA(image.Rect(0, 0, 18, 9))
	for i := 0; i < len(pxCoords); i += 2 {
		img.Set(pxCoords[i], pxCoords[i + 1], clr)
	}
	return ebiten.NewImageFromImage(img)
}
