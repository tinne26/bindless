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

type Title struct {
	renderer *etxt.Renderer
	horzPadWait int
	titleHorzPadding fixed.Int26_6
	tickCount uint32
	titlePadExpand bool
	titleOpacity uint8
	subOpacity uint8

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
	renderer.SetQuantizationMode(etxt.QuantizeNone)
	// TODO: ^ QuantizeVert should also work, this looks like a bug in etxt?

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
const untilSubText = 500
const startUpPreWait = 80

func (self *Title) Update(logCursorX, logCursorY int) error {
	const HorzPadWait = 60*6

	// detect mouse clicks to go to the next scene if enough ticks have passed
	if self.exitFadeout == 0 && self.tickCount > untilSubText && misc.MousePressed() {
		self.exitFadeout += 1
		sound.PlaySFX(sound.SfxAbility)
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
	if self.tickCount > untilSubText && self.subOpacity < 165 {
		self.subOpacity += 1
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

	// draw subtext if enough time has passed
	if self.tickCount >= untilSubText {
		// adjust padding and size for the subText
		sizer.SetHorzPadding(int(2*zoomLevel))
		self.renderer.SetSizePx(misc.ScaledFontSize(10, zoomLevel))

		// set the color and draw
		subOpacity := self.adjustedOpacity(self.subOpacity)
		self.renderer.SetColor(color.RGBA{240, 240, 240, subOpacity})
		self.renderer.Draw("left-click anywhere to start", x, y + int(32*zoomLevel))
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
