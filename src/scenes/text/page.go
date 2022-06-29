package text

import "image"

import "github.com/hajimehoshi/ebiten/v2"

import "bindless/src/misc"
import "bindless/src/misc/typewriter"
import "bindless/src/game/sceneitf"
import "bindless/src/sound"

type fadeType uint8
const (
	fadeIn   fadeType = 0
	fadeNone fadeType = 1
	fadeOut  fadeType = 2
)

type Page struct {
	writer *typewriter.Writer
	opacity uint8
	fade fadeType // fadeIn, fadeNone, fadeOut
	skipPressed bool
}

func New(ctx *misc.Context, key pageKey) (*Page, error) {
	return &Page {
		writer: typewriter.New(gameTexts[int(key)], ctx),
	}, nil
}

func (self *Page) Update(logCursorX, logCursorY int) error {
	if !misc.SkipKeyPressed() {
		self.skipPressed = false
	} else if !self.skipPressed {
		self.skipPressed = true
		if self.writer.ReachedEnd() {
			sound.PlaySFX(sound.SfxNope)
		} else {
			sound.PlaySFX(sound.SfxAbility)
			self.writer.SkipToEnd()
		}
	}
	self.writer.Tick()

	switch self.fade {
	case fadeIn:
		if self.opacity >= 252 { self.opacity = 255 } else { self.opacity += 3 }
		if self.opacity == 255 { self.fade = fadeNone }
	case fadeNone:
		if self.writer.ReachedEnd() && misc.MousePressed() {
			self.fade = fadeOut
			sound.PlaySFX(sound.SfxClick)
			sound.RequestFadeOut()
		}
	case fadeOut:
		if self.opacity <= 3 { self.opacity = 0 } else { self.opacity -= 3 }
	default:
		panic("invalid fade mode")
	}

	return nil
}

func (self *Page) Draw(screen *ebiten.Image) {}

func (self *Page) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	fontSize := misc.ScaledFontSize(10, zoomLevel)
	bounds := screen.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	horzMargin, vertMargin := w/8, h/10
	ox, oy := horzMargin, vertMargin
	fx, fy := w - horzMargin, h - vertMargin
	rect := image.Rect(bounds.Min.X + ox, bounds.Min.Y + oy, bounds.Min.X + fx, bounds.Min.Y + fy)
	self.writer.Draw(screen, fontSize, rect, self.opacity)
}

func (self *Page) Status() sceneitf.Status {
	if self.fade == fadeOut && self.opacity == 0 { return sceneitf.IsOver }
	return sceneitf.KeepAlive
}
