package palette

import "math/rand"
import "image/color"

var Background       = color.RGBA{ 47,  50,  58, 255}
var PolarityPositive = color.RGBA{159, 236, 232, 255}
var PolarityNegative = color.RGBA{250, 248, 132, 255} //251, 248, 125, 255}
var PolarityNeutral  = color.RGBA{200, 200, 200, 255}
var SceneWireframe   = color.RGBA{175,  82, 141, 255}
var AbilityDefault   = color.RGBA{240, 240, 240, 255}
var TileBottom       = color.RGBA{104,  81,  95, 255}
var Focus            = color.RGBA{255,  99, 112, 255}
var CircuitShadow    = color.RGBA{ 25,  25,  14,  72}

func SampleTileColor() color.RGBA {
	switch rand.Intn(3) {
	case 0: return color.RGBA{157, 106, 137, 255}
	case 1: return color.RGBA{178, 119, 155, 255}
	case 2: return color.RGBA{196, 131, 171, 255}
	default:
		panic("bad rng, bad!")
	}
}

func PreMultAlpha(clr color.RGBA) color.RGBA {
	clr.R = uint8((uint16(clr.R)*uint16(clr.A))/255)
	clr.G = uint8((uint16(clr.G)*uint16(clr.A))/255)
	clr.B = uint8((uint16(clr.B)*uint16(clr.A))/255)
	return clr
}

func Mix(over, back color.RGBA) color.RGBA {
	alpha := (uint32(over.A)*255 + uint32(back.A)*uint32(255 - over.A))/255
	if alpha > 255 { alpha = 255 }
	mixResult := color.RGBA{
		R: min((uint32(over.R)*255 + uint32(back.R)*uint32(255 - over.A))/255, alpha),
		G: min((uint32(over.G)*255 + uint32(back.G)*uint32(255 - over.A))/255, alpha),
		B: min((uint32(over.B)*255 + uint32(back.B)*uint32(255 - over.A))/255, alpha),
		A: uint8(alpha),
	}
	return mixResult
}

func min(a, b uint32) uint8 {
	if a <= b { return uint8(a) }
	return uint8(b)
}
