package level

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/art/palette"
import "github.com/tinne26/bindless/src/game/iso"

var twoPxImg *ebiten.Image
var surfaceOffscreen *ebiten.Image
func surface2img(surface iso.Map[struct{}]) *ebiten.Image {
	if twoPxImg == nil {
		twoPxImg = ebiten.NewImage(1, 2)
		twoPxImg.Fill(palette.TileBottom)
	}
	if surfaceOffscreen == nil {
		surfaceOffscreen = ebiten.NewImage(640, 360)
	} else {
		surfaceOffscreen.Clear()
	}

	surface.Each(
		func(col, row int16, _ struct{}) {
			// draw tile
			x, y := iso.TileCoords(col, row)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))
			opts.ColorM.ScaleWithColor(palette.SampleTileColor())
			surfaceOffscreen.DrawImage(graphics.TileMask, opts)

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
				surfaceOffscreen.DrawImage(graphics.TileBottomLeft, opts)
			}
			if !nextRowOccupied {
				opts.GeoM.Translate(18, 0)
				surfaceOffscreen.DrawImage(graphics.TileBottomRight, opts)
			}

			opts.GeoM.Reset()
			_, hasTileBelow := surface.Get(col - 1, row + 1)
			if !hasTileBelow {
				if prevColOccupied && !nextRowOccupied {
					opts.GeoM.Translate(float64(x + 17), float64(y + 18))
					surfaceOffscreen.DrawImage(twoPxImg, opts) // (17, 18-19)
				} else if !prevColOccupied && !nextRowOccupied {
					opts.GeoM.Translate(float64(x + 16), float64(y + 18))
					surfaceOffscreen.DrawImage(twoPxImg, opts) // (16, 18-19)
					opts.GeoM.Translate(1, 0)
					surfaceOffscreen.DrawImage(twoPxImg, opts) // (17, 18-19)
				} else {
					opts.GeoM.Translate(float64(x + 16), float64(y + 18))
					surfaceOffscreen.DrawImage(twoPxImg, opts) // (16, 18-19)
				}
			}
		})

	return surfaceOffscreen
}
