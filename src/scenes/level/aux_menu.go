package level

import "image/color"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/ui"
import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/art/palette"
import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/sound"

type AuxMenu struct {
	menu *ui.Menu
	active bool
	hovered bool
	mousePressed bool
	skipKeyPressed bool
}

func NewAuxMenu(menu *ui.Menu) *AuxMenu {
	return &AuxMenu { menu: menu }
}

func (self *AuxMenu) Update(logCursorX, logCursorY int) {
	// active case
	if self.active {
		skipKeyPressed := misc.SkipKeyPressed()
		if self.menu.AtRoot() && skipKeyPressed && !self.skipKeyPressed {
			self.skipKeyPressed = true
			sound.PlaySFX(sound.SfxClick)
			self.menu.Unselect()
			self.active  = false
			self.hovered = false
			return
		}

		self.menu.Update(logCursorX, logCursorY)
		self.skipKeyPressed = skipKeyPressed
		return
	}

	// detect hovering
	if logCursorX >= 640 - 10 - 21 && logCursorX <= 640 - 9 && logCursorY >= 360 - 10 - 16 && logCursorY <= 360 - 9 {
		if !self.hovered {
			self.hovered = true
			sound.PlaySFX(sound.SfxLoudNav)
		}
	} else {
		self.hovered = false
	}

	// detect clicks
	mousePressed := misc.MousePressed()
	if !mousePressed {
		self.mousePressed = false
	} else if !self.mousePressed {
		self.mousePressed = true
		if self.hovered {
			sound.PlaySFX(sound.SfxClick)
			self.active = true
		}
	}
}

func (self *AuxMenu) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(640 - 10 - 21, 360 - 9 - 16)
	screen.DrawImage(graphics.IconMenu, opts)
	if self.hovered && !self.active {
		opts.GeoM.Translate(0, -7)
		opts.ColorM.ScaleWithColor(palette.Background)
		screen.DrawImage(graphics.HudMenu, opts) // top shadow
		opts.GeoM.Translate(0, 2)
		screen.DrawImage(graphics.HudMenu, opts) // bottom shadow
		opts.GeoM.Translate(-1, -1)
		screen.DrawImage(graphics.HudMenu, opts) // left shadow
		opts.GeoM.Translate(2, 0)
		screen.DrawImage(graphics.HudMenu, opts) // right shadow
		opts.GeoM.Translate(-1, 0)
		opts.ColorM.Reset()
		screen.DrawImage(graphics.HudMenu, opts) // main text
	}
}

func (self *AuxMenu) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	if self.active {
		misc.DrawRect(screen, color.RGBA{0, 0, 0, 200})
		bounds := screen.Bounds()
		x, y := bounds.Dx()/2, bounds.Dy()/2
		self.menu.SetCenter(int(float64(x)/zoomLevel), int(float64(y)/zoomLevel))
		self.menu.SetBaseOpacity(164)
		self.menu.DrawHiRes(screen, zoomLevel)
	}
}