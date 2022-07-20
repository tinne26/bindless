package ui

import "image"
import "image/color"

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/tinne26/etxt"
import "github.com/tinne26/etxt/esizer"

import "github.com/tinne26/bindless/src/lang"
import "github.com/tinne26/bindless/src/sound"
import "github.com/tinne26/bindless/src/misc"

type Menu struct {
	logX int // logical center coordinate (x)
	logY int // logical y coordinate, may refer to center or top.
	         // by default it refers to center, and it's detected
				// through the renderer's vertical align

	renderer *etxt.Renderer
	baseOpacity uint8
	logHorzPadding int
	logOptSeparation int
	logTextSize int

	selDepth int
	tree *SubMenu

	prevFrameClick bool
	mousePressed bool
	skipKeyPressed bool
}

func NewMenu(ctx *misc.Context, font *etxt.Font, mainOptions []*lang.Text, handler func(string)) *Menu {
	renderer := etxt.NewStdRenderer()
	renderer.SetCacheHandler(ctx.FontCache.NewHandler())
	renderer.SetSizer(&esizer.HorzPaddingSizer{})
	renderer.SetFont(font)
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)

	menu := &Menu {
		logX: 640/2,
		logY: 360/2,
		renderer: renderer,
		baseOpacity: 136,
		logTextSize: 13,
		tree: newSubMenu(mainOptions, handler),
	}
	return menu
}

func (self *Menu) SetBaseOpacity(opacity uint8) {
	self.baseOpacity = opacity
}

func (self *Menu) SetLogHorzPadding(padding int) {
	self.logHorzPadding = padding
}

func (self *Menu) SetLogOptSeparation(separation int) {
	self.logOptSeparation = separation
}
func (self *Menu) SetLogFontSize(size int) {
	if size < 1 { panic("can't set logical font size < 1") }
	self.logTextSize = size
}

func (self *Menu) SetCenter(x, y int) {
	self.logX, self.logY = x, y
	self.renderer.SetAlign(etxt.YCenter, etxt.XCenter)
}

func (self *Menu) SetCenterTop(x, y int) {
	self.logX, self.logY = x, y
	self.renderer.SetAlign(etxt.Top, etxt.XCenter)
}

func (self *Menu) NavTo(option string) bool {
	subMenu := self.tree.NavDepth(self.selDepth)
	return subMenu.NavTo(option)
}

func (self *Menu) NavIn() bool {
	subMenu := self.tree.NavDepth(self.selDepth + 1)
	if subMenu == nil { return false }
	self.selDepth += 1
	return true
}

func (self *Menu) NavOut() {
	subMenu := self.tree.NavDepth(self.selDepth)
	subMenu.Unselect()
	if self.selDepth > 0 { self.selDepth -= 1 }
	return
}

func (self *Menu) NavRoot() {
	self.selDepth = 0
	self.tree.Unselect()
}

func (self *Menu) Unselect() {
	self.tree.NavDepth(self.selDepth).Unselect()
}

// Sets the submenu options for the current options level
func (self *Menu) SetSubOptions(option string, subOptions []*lang.Text, handler func(string)) bool {
	subMenu := self.tree.NavDepth(self.selDepth)
	return subMenu.SetSubOptions(option, subOptions, handler)
}

func (self *Menu) Update(logCursorX, logCursorY int) {
	// detect hover position
	subMenu := self.tree.NavDepth(self.selDepth)
	changed := self.updateHoverPosition(subMenu, logCursorX, logCursorY)
	if changed && !self.prevFrameClick {
		sound.PlaySFX(sound.SfxLoudNav)
	}
	self.prevFrameClick = false

	// handle cancel presses
	skipKeyPressed := misc.SkipKeyPressed()
	if !skipKeyPressed {
		self.skipKeyPressed = false
	} else if !self.skipKeyPressed {
		self.skipKeyPressed = true
		if self.selDepth > 0 {
			self.selDepth -= 1
			subMenu.Unselect()
			sound.PlaySFX(sound.SfxClick)
			self.prevFrameClick = true
		} else {
			sound.PlaySFX(sound.SfxNope)
		}
		return
	}

	// detect mouse clicks
	mousePressed := misc.MousePressed()
	if !mousePressed {
		self.mousePressed = false
	} else if !self.mousePressed {
		self.mousePressed = true
		if subMenu.selected != -1 {
			self.prevFrameClick = true
			subMenu.CallHandler()
		}
	}
}

func (self *Menu) updateHoverPosition(subMenu *SubMenu, logCursorX, logCursorY int) bool {
	y := self.logY
	cursor := image.Pt(logCursorX, logCursorY)
	self.renderer.SetSizePx(self.logTextSize)
	vertAlign, _ := self.renderer.GetAlign()
	var newSelection int = -1
	if vertAlign == etxt.YCenter {
		y -= (self.logOptSeparation*(len(subMenu.options) - 1))/2
		for i, opt := range subMenu.options {
			rect := self.renderer.SelectionRect(opt.Text.Get()).ImageRect()
			rect  = rect.Add(image.Pt(self.logX - rect.Dx()/2, y - rect.Dy()/2))
			if cursor.In(rect) {
				newSelection = i
				break
			}
			y += self.logOptSeparation
		}
	} else {
		for i, opt := range subMenu.options {
			rect := self.renderer.SelectionRect(opt.Text.Get()).ImageRect()
			rect  = rect.Add(image.Pt(self.logX - rect.Dx()/2, y))
			if cursor.In(rect) {
				newSelection = i
				break
			}
			y += self.logOptSeparation
		}
	}

	if newSelection != subMenu.selected {
		subMenu.selected = newSelection
		if newSelection != -1 {
			return true
		}
	}
	return false
}

func (self *Menu) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	// locate submenu to draw and compute misc sizes
	subMenu := self.tree.NavDepth(self.selDepth)
	optSeparation := int(float64(self.logOptSeparation)*zoomLevel)

	// configure renderer properties
	self.renderer.SetTarget(screen)
	if self.logHorzPadding != 0 {
		sizer := self.renderer.GetSizer().(*esizer.HorzPaddingSizer)
		sizer.SetHorzPadding(int(3*zoomLevel))
	}
	self.renderer.SetSizePx(misc.ScaledFontSize(float64(self.logTextSize), zoomLevel))
	self.renderer.SetHorzAlign(etxt.XCenter)

	// find drawing position, which depends on align
	b := screen.Bounds()
	x, y := b.Min.X + int(zoomLevel*float64(self.logX)), b.Min.Y + int(zoomLevel*float64(self.logY))
	vertAlign, _ := self.renderer.GetAlign()
	if vertAlign == etxt.YCenter {
		y -= (optSeparation*(len(subMenu.options) - 1))/2
	}

	// draw submenu
	for i, option := range subMenu.options {
		if i == subMenu.selected {
			self.renderer.SetColor(color.RGBA{240, 240, 240, 255})
		} else {
			self.renderer.SetColor(color.RGBA{240, 240, 240, self.baseOpacity})
		}
		self.renderer.Draw(option.Text.Get(), x, y)
		y += optSeparation
	}
}
