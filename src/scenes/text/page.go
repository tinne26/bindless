package text

import "fmt"
import "image"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/misc/typewriter"
import "github.com/tinne26/bindless/src/game/sceneitf"
import "github.com/tinne26/bindless/src/sound"
import "github.com/tinne26/bindless/src/ui"

type fadeType uint8
const (
	fadeIn   fadeType = 0
	fadeNone fadeType = 1
	fadeOut  fadeType = 2
)

type Page struct {
	writer *typewriter.Writer
	choices *ui.HorzChoice
	choicesOpacity uint8
	opacity uint8
	fade fadeType // fadeIn, fadeNone, fadeOut
	skipPressed bool
	endStatus sceneitf.Status
}

func New(ctx *misc.Context, key pageKey) (*Page, error) {
	page := &Page {
		writer: typewriter.New(gameTexts[int(key)].Get(), ctx),
	}
	choices := gameChoices[int(key)]
	if choices != nil {
		coda := ctx.FontLib.GetFont("Coda Regular")
		if coda == nil { return nil, fmt.Errorf("missing 'Coda' font") }
		page.choices = ui.NewHorzChoice(ctx, coda)
		page.choices.SetLogOptSeparation(4)
		for _, choice := range choices {
			switch choice.Text.English() {
			case "[ Play the tutorial ]", "[ Hmm.. can I repeat the tutorial? ]":
				choice.Handler = page.fnToTutorialHandler
			case "[ Skip the tutorial ]", "[ LET'S GET TO IT! ]":
				choice.Handler = page.fnToStoryHandler
			}
			page.choices.AddHChoice(choice)
		}
	}

	return page, nil
}

func (self *Page) Update(logCursorX, logCursorY int) error {
	if !misc.SkipKeyPressed() {
		self.skipPressed = false
	} else if !self.skipPressed {
		self.skipPressed = true
		if self.writer.ReachedEnd() {
			sound.SfxNope.Play()
		} else {
			sound.SfxAbility.Play()
			self.writer.SkipToEnd()
		}
	}
	self.writer.Tick()

	if self.writer.ReachedEnd() && self.choices != nil {
		if self.choicesOpacity >= 252 {
			self.choicesOpacity = 255
		} else {
			self.choicesOpacity += 3
		}
		
		if self.choicesOpacity > 80 {
			self.choices.Update(logCursorX, logCursorY)
		}
	}

	switch self.fade {
	case fadeIn:
		if self.opacity >= 252 { self.opacity = 255 } else { self.opacity += 3 }
		if self.opacity == 255 { self.fade = fadeNone }
	case fadeNone:
		if self.writer.ReachedEnd() && self.choices == nil && misc.MousePressed() {
			self.fade = fadeOut
			self.endStatus = sceneitf.IsOverNext
			sound.SfxClick.Play()
			//sound.RequestFadeOut()
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
	_, endY := self.writer.Draw(screen, fontSize, rect, self.opacity)
	if self.writer.ReachedEnd() && self.choices != nil {
		var opacity uint8
		if self.fade == fadeOut {
			self.choices.Unfocus()
			opacity = self.opacity
			if self.choicesOpacity < opacity {
				opacity = self.choicesOpacity
			}
		} else {
			opacity = self.choicesOpacity
		}
		
		self.choices.SetBaseOpacity(uint8((float64(opacity)/255)*128))		
		self.choices.SetTopLeft(int(float64(horzMargin)/zoomLevel), int(float64(endY - bounds.Min.Y)/zoomLevel) + 13)
		self.choices.DrawHiRes(screen, zoomLevel)
	}
}

func (self *Page) Status() sceneitf.Status {
	if self.fade == fadeOut && self.opacity == 0 {
		return self.endStatus
	}
	return sceneitf.KeepAlive
}

func (self *Page) fnToStoryHandler(_ string) {
	self.fade = fadeOut
	self.endStatus = sceneitf.ToStory
	sound.SfxClick.Play()
}

func (self *Page) fnToTutorialHandler(_ string) {
	self.fade = fadeOut
	self.endStatus = sceneitf.ToTutorial
	sound.SfxClick.Play()
}