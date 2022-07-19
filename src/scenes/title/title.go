package title

import "fmt"
import "image/color"

import "golang.org/x/image/math/fixed"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/etxt"
import "github.com/tinne26/etxt/esizer"

import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/game/sceneitf"
import "github.com/tinne26/bindless/src/sound"
import "github.com/tinne26/bindless/src/lang"

type Title struct {
	renderer *etxt.Renderer
	horzPadWait int
	titleHorzPadding fixed.Int26_6
	tickCount uint32
	titlePadExpand bool
	titleOpacity uint8
	optsOpacity uint8
	mousePressed bool
	skipKeyPressed bool
	inLangSubmenu bool
	optHoverIndex uint8
	exitFadeout uint8
}

func New(ctx *misc.Context) (*Title, error) {
	// create text renderer
	renderer := etxt.NewStdRenderer()
	renderer.SetCacheHandler(ctx.FontCache.NewHandler())
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)

	// add a custom sizer to animate the title text and
	// make the title renderer horizontally unquantized
	renderer.SetSizer(&esizer.HorzPaddingSizer{})
	renderer.SetQuantizationMode(etxt.QuantizeVert)
	// ^ either etxt.QuantizeVert or etxt.QuantizeNone will have
	//   the same effect here. It's only important to avoid
	//   horizontal quantization to ensure smooth animation.

	// load fonts from context and set them
	coda := ctx.FontLib.GetFont("Coda Regular")
	if coda == nil { return nil, fmt.Errorf("missing 'Coda' font") }
	renderer.SetFont(coda)

	return &Title {
		renderer: renderer,
		titlePadExpand: true,
		titleHorzPadding: baseTitleHorzFract,
	}, nil
}

// constants related to text animations
const baseTitleHorzFract = 560  // fixed.Int26_6 (divide by 64)
const maxTitleHorzFract  = 1230 // fixed.Int26_6
const untilOptsText = 500
const untilCanClick = 540
const startUpPreWait = 80
const optLogicalSeparation = 26

func (self *Title) Update(logCursorX, logCursorY int) error {
	const HorzPadWait = 60*6

	// update mouse interaction
	if self.exitFadeout == 0 && self.tickCount > untilCanClick {
		// detect hover position
		hoverIndex := uint8(0)
		if logCursorX > 270 && logCursorX < 640 - 270 {
			startY := 177
			for i := 0; i < 4; i++ {
				y := startY + i*optLogicalSeparation
				if logCursorY > y - 9 && logCursorY < y + 9 {
					hoverIndex = uint8(i + 1)
					break
				}
			}
		}
		if !self.inLangSubmenu && hoverIndex == 4 { hoverIndex = 0 }
		if hoverIndex != self.optHoverIndex {
			self.optHoverIndex = hoverIndex
			if hoverIndex != 0 {
				sound.PlaySFX(sound.SfxLoudNav)
			}
		}

		// handle ESC/TAB key presses if in lang submenu to go back
		skipKeyPressed := misc.SkipKeyPressed()
		if !skipKeyPressed {
			self.skipKeyPressed = false
		} else if !self.skipKeyPressed {
			self.skipKeyPressed = true
			if self.inLangSubmenu {
				self.inLangSubmenu = false
				if self.optHoverIndex == 4 { self.optHoverIndex = 0 }
				sound.PlaySFX(sound.SfxClick)
			}
		}

		// detect mouse clicks
		mousePressed := misc.MousePressed()
		if !mousePressed {
			self.mousePressed = false
		} else if !self.mousePressed {
			self.mousePressed = true
			if self.inLangSubmenu {
				switch self.optHoverIndex {
				case 1: // english
					lang.Set(lang.EN)
					sound.PlaySFX(sound.SfxClick)
					self.inLangSubmenu = false
				case 2: // spanish
					lang.Set(lang.ES)
					sound.PlaySFX(sound.SfxClick)
					self.inLangSubmenu = false
				case 3: // catalan
					lang.Set(lang.CA)
					sound.PlaySFX(sound.SfxClick)
					self.inLangSubmenu = false
				case 4: // back
					sound.PlaySFX(sound.SfxClick)
					self.inLangSubmenu = false
					self.optHoverIndex = 0
				}
			} else {
				switch self.optHoverIndex {
				case 1: // start game
					sound.PlaySFX(sound.SfxAbility)
					self.exitFadeout += 1
					self.optHoverIndex = 0
				case 2: // language
					self.inLangSubmenu = true
					sound.PlaySFX(sound.SfxClick)
				case 3: // fullscreen switch
					sound.PlaySFX(sound.SfxClick)
					ebiten.SetFullscreen(!ebiten.IsFullscreen())
				}
			}
		}
	}

	// if performing exit fadeout, increase the counter
	if self.exitFadeout > 0 && self.exitFadeout < 255 {
		self.exitFadeout += 2
	}

	// increase tick count
	self.tickCount += 1
	if self.tickCount < startUpPreWait { return nil }

	// dimiss odd ticks
	if self.tickCount & 0x1 == 0x1 { return nil }

	// increase text opacities
	if self.titleOpacity < 220 { self.titleOpacity += 1 }
	if self.tickCount > untilOptsText && self.optsOpacity < 136 {
		self.optsOpacity += 1
	}

	// temporary wait when title is fully extended or contracted
	if self.horzPadWait > 0 { self.horzPadWait -= 1 }

	// update title expansion / contraction
	if self.horzPadWait == 0 {
		if self.titlePadExpand {
			// expanding
			self.titleHorzPadding += 1
			if self.titleHorzPadding >= maxTitleHorzFract {
				self.titleHorzPadding = maxTitleHorzFract
				self.titlePadExpand = false
				self.horzPadWait = HorzPadWait
			}
		} else { // self.titlePadExpand == false
			// contracting
			self.titleHorzPadding -= 1
			if self.titleHorzPadding <= baseTitleHorzFract {
				self.titleHorzPadding = baseTitleHorzFract
				self.titlePadExpand = true
				self.horzPadWait = HorzPadWait
			}
		}
	}

	return nil
}

func (self *Title) Draw(screen *ebiten.Image) {}

func (self *Title) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	if self.tickCount < startUpPreWait { return }
	//if self.Status() == IsOver { return }

	// adjust draw target, padding and size
	self.renderer.SetTarget(screen)
	sizer := self.renderer.GetSizer().(*esizer.HorzPaddingSizer)
	sizer.SetHorzPaddingFract(fixed.Int26_6(float64(self.titleHorzPadding)*zoomLevel))
	self.renderer.SetSizePx(misc.ScaledFontSize(34, zoomLevel))

	// set title opacity
	titleOpacity := uint8(220)
	titleOpacity = self.adjustedOpacity(self.titleOpacity)
	self.renderer.SetColor(color.RGBA{255, 255, 255, titleOpacity})

	// determine drawing position
	bounds := screen.Bounds()
	x := bounds.Min.X + bounds.Dx()/2
	height := bounds.Dy()
	y := bounds.Min.Y + height/2 - height/8
	self.renderer.Draw("BINDLESS", x, y)

	// draw options if enough time has passed
	if self.tickCount >= untilOptsText {
		// adjust padding and size for the options
		sizer.SetHorzPadding(int(3*zoomLevel))
		self.renderer.SetSizePx(misc.ScaledFontSize(13, zoomLevel))
		optsOpacity := self.adjustedOpacity(self.optsOpacity)

		// draw options
		y += int(42*zoomLevel)
		optSeparation := int(optLogicalSeparation*zoomLevel)
		if self.inLangSubmenu {
			self.setOptColor(1, optsOpacity)
			self.renderer.Draw("English", x, y)
			y += optSeparation
			self.setOptColor(2, optsOpacity)
			self.renderer.Draw("Español", x, y)
			y += optSeparation
			self.setOptColor(3, optsOpacity)
			self.renderer.Draw("Català", x, y)
			y += optSeparation
			self.setOptColor(4, optsOpacity)
			self.renderer.Draw("< " + lang.Tr("Back") + " >", x, y)
		} else {
			self.setOptColor(1, optsOpacity)
			self.renderer.Draw(lang.Tr("Start Game"), x, y)
			y += optSeparation
			self.setOptColor(2, optsOpacity)
			self.renderer.Draw(lang.Tr("Language"), x, y)
			y += optSeparation
			self.setOptColor(3, optsOpacity)
			self.renderer.Draw(lang.Tr("Fullscreen"), x, y)
		}
	}
}

func (self *Title) setOptColor(index uint8, optsOpacity uint8) {
	if index == self.optHoverIndex {
		self.renderer.SetColor(color.RGBA{240, 240, 240, 255})
	} else {
		self.renderer.SetColor(color.RGBA{240, 240, 240, optsOpacity})
	}
}

func (self *Title) Status() sceneitf.Status {
	if self.exitFadeout == 255 { return sceneitf.IsOver }
	return sceneitf.KeepAlive
}

// during the exit fadeout, opacities need to be overridden
func (self *Title) adjustedOpacity(opacity uint8) uint8 {
	if self.exitFadeout == 0 { return opacity }
	if self.exitFadeout >= opacity { return 0 }
	return opacity - self.exitFadeout
}
