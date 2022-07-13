package dev

import "github.com/hajimehoshi/ebiten/v2"

type RaisingMagnet struct {
	polarity PolarityType
	x int
	y int
}

func (self *RaisingMagnet) Update() { self.y -= 1 }
func (self *RaisingMagnet) Y() int { return self.y }
func (self *RaisingMagnet) Draw(screen *ebiten.Image) {
	drawSmallMagnetAt(screen, self.x, self.y, self.polarity, self.polarity.Color(), false)
}
