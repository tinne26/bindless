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
var AbilitySelected  = color.RGBA{255,  99, 112, 255}
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
