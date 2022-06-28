package level

import "github.com/hajimehoshi/ebiten/v2"

import "bindless/src/art/graphics"
import "bindless/src/art/palette"
import "bindless/src/game/iso"


func surface2img(surface iso.Map[struct{}]) *ebiten.Image {
	img := ebiten.NewImage(640, 360)

	surface.Each(
		func(col, row int16, _ struct{}) {
			// draw tile
			x, y := iso.TileCoords(col, row)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))
			opts.ColorM.ScaleWithColor(palette.SampleTileColor())
			img.DrawImage(graphics.TileMask, opts)

			// to check for bottom borders to draw, we follow this logic:
			// - if tile at col-1 doesn't exist, draw main left part
			// - if tile at row+1 doesn't exist, draw main right part
			// - if tile at (col-1, row+1) doesn't exist, then...
			// 	- if has col-1 and not row+1, single 2 pix right
			// 	- if doesn't have any, full square
			// 	- else, single 2 pix left
			_, prevColOccupied := surface.Get(col - 1, row)
			_, nextRowOccupied := surface.Get(col, row + 1)
			opts.ColorM.Reset()
			opts.GeoM.Translate(0, 10)
			if !prevColOccupied {
				img.DrawImage(graphics.TileBottomLeft, opts)
			}
			if !nextRowOccupied {
				opts.GeoM.Translate(18, 0)
				img.DrawImage(graphics.TileBottomRight, opts)
			}

			_, hasTileBelow := surface.Get(col - 1, row + 1)
			if !hasTileBelow {
				if prevColOccupied && !nextRowOccupied {
					img.Set(x + 17, y + 18, palette.TileBottom)
					img.Set(x + 17, y + 19, palette.TileBottom)
				} else if !prevColOccupied && !nextRowOccupied {
					img.Set(x + 16, y + 18, palette.TileBottom)
					img.Set(x + 16, y + 19, palette.TileBottom)
					img.Set(x + 17, y + 18, palette.TileBottom)
					img.Set(x + 17, y + 19, palette.TileBottom)
				} else {
					img.Set(x + 16, y + 18, palette.TileBottom)
					img.Set(x + 16, y + 19, palette.TileBottom)
				}
			}
		})

	return img
}
