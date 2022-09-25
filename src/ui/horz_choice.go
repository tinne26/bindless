package ui

import "image"
import "image/color"

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/tinne26/etxt"
import "github.com/tinne26/etxt/esizer"

import "github.com/tinne26/bindless/src/lang"
import "github.com/tinne26/bindless/src/sound"
import "github.com/tinne26/bindless/src/misc"

type HChoice struct {
	Text *lang.Text
	Handler func(string) // option click handler, or nil if disabled
}

type HorzChoice struct {
	leftX int
	topY  int

	renderer *etxt.Renderer
	baseOpacity uint8
	logHorzPadding int
	logOptSeparation int
	logTextSize int

	prevFrameClick bool
	mousePressed bool
	unfocused bool
	choices []*HChoice
	hoverIndex int
}

func NewHorzChoice(ctx *misc.Context, font *etxt.Font) *HorzChoice {
	renderer := etxt.NewStdRenderer()
	renderer.SetCacheHandler(ctx.FontCache.NewHandler())
	renderer.SetSizer(&esizer.HorzPaddingSizer{})
	renderer.SetFont(font)
	renderer.SetAlign(etxt.Top, etxt.Left)

	choice := &HorzChoice {
		leftX: 0,
		topY: 0,
		renderer: renderer,
		baseOpacity: 128,
		logTextSize: 10,
		choices: make([]*HChoice, 0, 4),
		hoverIndex: -1,
	}
	return choice
}

func (self *HorzChoice) AddChoice(text *lang.Text, handler func(string)) {
	opt := &HChoice {
		Text: text,
		Handler: handler,
	}
	self.choices = append(self.choices, opt)
}

func (self *HorzChoice) AddHChoice(opt *HChoice) {
	self.choices = append(self.choices, opt)
}

func (self *HorzChoice) SetBaseOpacity(opacity uint8) {
	self.baseOpacity = opacity
}

func (self *HorzChoice) Unfocus() {
	self.unfocused = true
}

func (self *HorzChoice) SetLogHorzPadding(padding int) {
	self.logHorzPadding = padding
}

func (self *HorzChoice) SetLogOptSeparation(separation int) {
	self.logOptSeparation = separation
}

func (self *HorzChoice) SetLogFontSize(size int) {
	if size < 1 { panic("can't set logical font size < 1") }
	self.logTextSize = size
}

func (self *HorzChoice) SetTopLeft(x, y int) {
	self.leftX, self.topY = x, y
}

func (self *HorzChoice) Update(logCursorX, logCursorY int) {
	if self.unfocused {
		self.hoverIndex = -1
		return
	}

	// detect hover position
	changed := self.updateHoverPosition(logCursorX, logCursorY)
	if changed && !self.prevFrameClick {
		sound.SfxNav.Play()
	}
	self.prevFrameClick = false

	// detect mouse clicks
	mousePressed := misc.MousePressed()
	if !mousePressed {
		self.mousePressed = false
	} else if !self.mousePressed {
		self.mousePressed = true
		if self.hoverIndex != -1 {
			self.prevFrameClick = true
			choice := self.choices[self.hoverIndex]
			if choice.Handler == nil {
				sound.SfxNope.Play()
			} else {
				choice.Handler(choice.Text.English())
			}
		}
	}
}

func (self *HorzChoice) updateHoverPosition(logCursorX, logCursorY int) bool {
	x := self.leftX
	cursor := image.Pt(logCursorX, logCursorY)
	self.renderer.SetSizePx(self.logTextSize)
	var newSelection int = -1
	
	for i, opt := range self.choices {
		rect := self.renderer.SelectionRect(opt.Text.Get())
		imgRect := rect.ImageRect().Add(image.Pt(x, self.topY))
		if cursor.In(imgRect) {
			newSelection = i
			break
		}
		x += rect.Width.Ceil() + self.logOptSeparation
	}

	if newSelection != self.hoverIndex {
		self.hoverIndex = newSelection
		if newSelection != -1 {
			return true
		}
	}
	return false
}

func (self *HorzChoice) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	// configure renderer properties
	self.renderer.SetTarget(screen)
	if self.logHorzPadding != 0 {
		sizer := self.renderer.GetSizer().(*esizer.HorzPaddingSizer)
		sizer.SetHorzPadding(int(float64(self.logHorzPadding)*zoomLevel))
	}
	self.renderer.SetSizePx(misc.ScaledFontSize(float64(self.logTextSize), zoomLevel))

	// find drawing position
	b := screen.Bounds()
	x, y := b.Min.X + int(zoomLevel*float64(self.leftX)), b.Min.Y + int(zoomLevel*float64(self.topY))
	optSeparation := int(float64(self.logOptSeparation)*zoomLevel)

	// draw submenu
	for i, option := range self.choices {
		if i != self.hoverIndex || option.Handler == nil {
			self.renderer.SetColor(color.RGBA{240, 240, 240, self.baseOpacity})
		} else {
			self.renderer.SetColor(color.RGBA{240, 240, 240, 255})
		}
		text := option.Text.Get()
		self.renderer.Draw(text, x, y)
		rect := self.renderer.SelectionRect(text)
		x += rect.Width.Ceil() + optSeparation
	}
}
