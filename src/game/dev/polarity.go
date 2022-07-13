package dev

import "image/color"

import "github.com/tinne26/bindless/src/art/palette"

type PolarityType uint8
const (
	PolarityNeutral  PolarityType = 0
	PolarityPositive PolarityType = 1
	PolarityNegative PolarityType = 2
	PolarityHack PolarityType = 66
)

func (self PolarityType) Color() color.RGBA {
	switch self {
	case PolarityNeutral : return palette.PolarityNeutral
	case PolarityPositive: return palette.PolarityPositive
	case PolarityNegative: return palette.PolarityNegative
	default:
		return color.RGBA{255, 0, 255, 255}
	}
}

var polNeutralFunc  = func() (PolarityType, color.RGBA) { return PolarityNeutral , palette.PolarityNeutral  }
var polPositiveFunc = func() (PolarityType, color.RGBA) { return PolarityPositive, palette.PolarityPositive }
var polNegativeFunc = func() (PolarityType, color.RGBA) { return PolarityNegative, palette.PolarityNegative }
func (self PolarityType) AsFunc() func() (PolarityType, color.RGBA) {
	switch self {
	case PolarityNeutral : return polNeutralFunc
	case PolarityPositive: return polPositiveFunc
	case PolarityNegative: return polNegativeFunc
	default:
		panic("nono")
	}
}

type Polarized interface {
	Polarity() PolarityType
}
