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
import "github.com/tinne26/bindless/src/ui"

type Title struct {
	renderer *etxt.Renderer
	menu *ui.Menu
	horzPadWait int
	titleHorzPadding fixed.Int26_6
	tickCount uint32
	titlePadExpand bool
	titleOpacity uint8
	optsOpacity uint8
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

	// create scene
	title := &Title {
		renderer: renderer,
		titlePadExpand: true,
		titleHorzPadding: baseTitleHorzFract,
	}

	// create menu
	optsRoot := []*lang.Text{
		lang.NewText("Start Game", "Empezar", "Començar"),
		lang.NewText("Language", "Idioma", "Idioma"),
		lang.NewText("Fullscreen", "Resolución", "Resolució"),
	}
	optsLang := []*lang.Text{
		lang.NewText("English", "English", "English"),
		lang.NewText("Español", "Español", "Español"),
		lang.NewText("Català", "Català", "Català"),
		lang.NewText("-- Back --", "-- Atrás --", "-- Enrere --"),
	}
	menu := ui.NewMenu(ctx, coda,  optsRoot, title.menuRootHandler)
	menu.SetSubOptions("Language", optsLang, title.menuLangHandler)
	menu.SetLogFontSize(13)
	menu.SetLogOptSeparation(26)
	menu.SetLogHorzPadding(3)

	title.menu = menu
	return title, nil
}

// constants related to text animations
const baseTitleHorzFract = 560  // fixed.Int26_6 (divide by 64)
const maxTitleHorzFract  = 1230 // fixed.Int26_6
const untilOptsText = 500
const untilCanClick = 540
const startUpPreWait = 80

func (self *Title) Update(logCursorX, logCursorY int) error {
	const HorzPadWait = 60*6

	// update interactive menu
	if self.exitFadeout == 0 && self.tickCount > untilCanClick {
		self.menu.Update(logCursorX, logCursorY)
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

	// draw menu if enough time has passed
	if self.tickCount >= untilOptsText {
		optsOpacity := self.adjustedOpacity(self.optsOpacity)
		x, y := bounds.Dx()/2, int(float64(bounds.Dy())*0.4722)
		self.menu.SetCenterTop(int(float64(x)/zoomLevel), int(float64(y)/zoomLevel))
		self.menu.SetBaseOpacity(optsOpacity)
		self.menu.DrawHiRes(screen, zoomLevel)
	}
}

func (self *Title) Status() sceneitf.Status {
	if self.exitFadeout == 255 { return sceneitf.IsOverNext }
	return sceneitf.KeepAlive
}

// during the exit fadeout, opacities need to be overridden
func (self *Title) adjustedOpacity(opacity uint8) uint8 {
	if self.exitFadeout == 0 { return opacity }
	if self.exitFadeout >= opacity { return 0 }
	return opacity - self.exitFadeout
}

func (self *Title) menuRootHandler(option string) {
	switch option {
	case "Start Game":
		sound.PlaySFX(sound.SfxAbility)
		self.exitFadeout += 1
		self.menu.Unselect()
	case "Language": // language
		sound.PlaySFX(sound.SfxClick)
		self.menu.NavIn()
	case "Fullscreen": // fullscreen switch
		sound.PlaySFX(sound.SfxClick)
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	default:
		panic(option)
	}
}

func (self *Title) menuLangHandler(option string) {
	self.menu.NavOut()
	switch option {
	case "English":
		lang.Set(lang.EN)
		sound.PlaySFX(sound.SfxClick)
	case "Español":
		lang.Set(lang.ES)
		sound.PlaySFX(sound.SfxClick)
	case "Català":
		lang.Set(lang.CA)
		sound.PlaySFX(sound.SfxClick)
	case "-- Back --":
		sound.PlaySFX(sound.SfxClick)
	default:
		panic(option)
	}
}
