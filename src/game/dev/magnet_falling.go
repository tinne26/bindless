package dev

import "image"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/shaders"

type FallingMagnet struct {
	inSpectre bool
	polarity PolarityType
	y float64
	speed float64
	origX int
	origY int
}

func (self *FallingMagnet) LogicalY() int { return int(self.y) }

// if returns true, can delete already for being out of screen
func (self *FallingMagnet) FallUpdate() bool {
	self.y += self.speed
	if self.speed < 2.3 { self.speed += 0.04 }
	return self.y >= 360
}

var offscreenMagnetCanvas *ebiten.Image
var copyCanvas *ebiten.Image
func (self *FallingMagnet) Draw(screen *ebiten.Image, _ float64) {
	// setup offscreens
	if copyCanvas == nil {
		copyCanvas = ebiten.NewImage(16, 16)
	} else {
		copyCanvas.Clear()
	}
	if offscreenMagnetCanvas == nil {
		offscreenMagnetCanvas = ebiten.NewImage(16, 16)
	} else {
		offscreenMagnetCanvas.Clear()
	}

	// draw magnet and apply shaders to cut the proper area
	drawSmallMagnetAt(offscreenMagnetCanvas, 0, 0, self.polarity, self.inSpectre)
	baseX, baseY := int(self.origX), int(self.y)
	screenRegion := screen.SubImage(image.Rect(baseX, baseY, baseX + 16, baseY + 16)).(*ebiten.Image)
	_, h := screenRegion.Size()
	imgOpts := &ebiten.DrawImageOptions{}
	imgOpts.GeoM.Translate(0, float64(16 - h))
	copyCanvas.DrawImage(screenRegion, imgOpts)
	opts := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"CutY": float32(self.origY + 16), // TODO: adjust value when testing floating tiles
			"RefX": float32(self.origX),
		},
	}
	opts.Images[0] = offscreenMagnetCanvas
	opts.Images[1] = copyCanvas
	opts.GeoM.Translate(float64(baseX), float64(baseY))
	screenRegion.DrawRectShader(16, 16, shaders.FallMagnetCut, opts)
}
