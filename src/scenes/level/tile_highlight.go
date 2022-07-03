package level

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/game/iso"

// The tile highlight is used to give visual feedback to the player
// about which tile is currently being selected / hovered.
// Internally it has two variations: one where a full tile is shown,
// and another where a tile partially cut is shown in order to make
// it look better when there's a magnet in the middle of the tile.

type tileHighlight struct {
	col int16
	row int16
	active bool
	cutting bool
}

func (self *tileHighlight) LogicalY() int {
	y := iso.YCoord(self.col, self.row)
	if self.cutting {
		return y + 1 // to draw floating cut highlight above magnet
	} else {
		return y - 1 // to draw normal floating highlight below magnet
	}
}

func (self *tileHighlight) Draw(screen *ebiten.Image, _ float64) {
	if !self.active { return }
	x, y := iso.TileCoords(self.col, self.row)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y - 2))
	opts.ColorM.Scale(1.0, 1.0, 1.0, 0.5)
	if self.cutting {
		screen.DrawImage(graphics.CutTileHighlight, opts)
	} else {
		screen.DrawImage(graphics.TileMask, opts)
	}
}
