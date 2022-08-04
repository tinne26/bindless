package dev

import "image"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/art/palette"

// TODO: maybe add effect type, to allow using targets for hints too..?

type Target struct {
	positions []image.Point
	posIndex int
	looping bool
	opacityFalling bool
	opacity uint8
	waitCounter uint8
}

const targetLowOpacityLimit = 226
func NewTarget(initCol, initRow int16, looping bool) *Target {
	x, y := iso.TileCoords(initCol, initRow)
	return &Target { positions: []image.Point{ image.Pt(x, y) }, looping: looping, opacity: targetLowOpacityLimit }
}

func (self *Target) AddPos(col, row int16) {
	x, y := iso.TileCoords(col, row)
	self.positions = append(self.positions, image.Pt(x, y))
}

func (self *Target) Draw(screen *ebiten.Image) {
	if self.posIndex == -1 { return }
	opts := &ebiten.DrawImageOptions{}
	position := self.positions[self.posIndex]
	opts.GeoM.Translate(float64(position.X) + 9, float64(position.Y) + 3)
	screen.DrawImage(graphics.TargetShadow, opts)
	clr := palette.Focus
	clr.A = self.opacity
	opts.ColorM.ScaleWithColor(clr)
	opts.GeoM.Translate(0, 1)
	screen.DrawImage(graphics.TargetShape, opts)
}

func (self *Target) Update() {
	if self.posIndex == -1 { return }

	// apply wait
	if self.waitCounter > 0 {
		self.waitCounter -= 1
		return
	} else {
		self.waitCounter = 2
	}

	// modify opacity
	if self.opacityFalling {
		self.opacity -= 1
		if self.opacity == targetLowOpacityLimit {
			self.opacityFalling = false
		}
	} else {
		self.opacity += 1
		if self.opacity == 255 {
			self.opacityFalling = true
		}
	}
}

func (self *Target) GetColRow() (int16, int16) {
	if self.posIndex == -1 { return -1, -1 }
	position := self.positions[self.posIndex]
	return iso.CoordsToTile(position.X + 8, position.Y + 4)
}

func (self *Target) Move() {
	self.posIndex += 1
	if self.posIndex >= len(self.positions) {
		if self.looping {
			self.posIndex = 0
		} else {
			self.posIndex = -1
		}
	}
}
