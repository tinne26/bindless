package level

import "image"
import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/game/iso"
import "github.com/tinne26/bindless/src/art/palette"

type AbilityExecCue struct { x, y int }

func NewAbilityExecCue(col, row int16) AbilityExecCue {
	x, y := iso.TileCoords(col, row)
	return AbilityExecCue{ x, y }
}

func (self *AbilityExecCue) Draw(screen *ebiten.Image, cycle float64) {
	xOffset := 4
	x, y := self.x + 4, self.y + 10
	w := 34 - xOffset
	screen.SubImage(image.Rect(x, y, x + w, y + 3)).(*ebiten.Image).Fill(palette.Background)
	x += 1
	y += 1
	w -= 2
	screen.SubImage(image.Rect(x, y, x + w, y + 1)).(*ebiten.Image).Fill(color.RGBA{180, 180, 180, 255})
	w = int(cycle*float64(w))
	if w == 0 { return }
	screen.SubImage(image.Rect(x, y, x + w, y + 1)).(*ebiten.Image).Fill(palette.Focus)
}
