package shaders

import _ "embed"

import "github.com/hajimehoshi/ebiten/v2"

//go:embed fall_magnet_cut.kage
var fallMagnetCutSrc []byte

var FallMagnetCut *ebiten.Shader

func Load() {
	var err error
	FallMagnetCut, err = ebiten.NewShader(fallMagnetCutSrc)
	if err != nil { panic(err) }
}
