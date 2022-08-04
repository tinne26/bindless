package typewriter

import "image"
import "image/color"
import "math/rand"

import "github.com/hajimehoshi/ebiten/v2"
import "golang.org/x/image/math/fixed"
import "github.com/tinne26/etxt"

import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/sound"

const waitStd  = 2.5
const pauseStd = 3.1
const intermittentCycleTicks = 120

// TODO: backtracking would be cool, but it's tricky and I don't have
//       the time to spare. we could use a tree with branches or
//       something like that?

// code similar to https://github.com/tinne26/etxt/blob/main/examples/ebiten/typewriter/main.go
// but with simpler, non-nestable formatting rules:
// - \x01: use current x as next line start x
// - \x02: reset next line start x to original value
// - \x03: pause *6
// - \x04: pause *10
// - \x05: pause *13
// - \x07: set text hue to main color
// - \x08: set text hue to disabled color
// - \x09: set text hue to name color
// - \x0A: line break
// - \x0B: 50% line height shift, to be used after normal line breaks
// - \x0C: shake off
// - \x0D: shake on
// - \x0E: shake on soft
// - \x10: no-wait period .
// - \x11: intermittent off
// - \x12: intermittent on

type Writer struct {
	renderer *etxt.Renderer
	text string
	index int
	wait float64
	intermittent int
}

func New(text string, ctx *misc.Context) *Writer {
	renderer := etxt.NewStdRenderer()
	renderer.SetCacheHandler(ctx.FontCache.NewHandler())
	renderer.SetFont(ctx.FontLib.GetFont("Coda Regular"))
	renderer.SetAlign(etxt.Top, etxt.Left)
	renderer.SetQuantizationMode(etxt.QuantizeNone)

	return &Writer {
		renderer: renderer,
		text: text,
		wait: waitStd,
	}
}

func (self *Writer) ReachedEnd() bool {
	return self.index == len(self.text)
}

func (self *Writer) SkipToEnd() {
	self.wait = 0
	self.index = len(self.text)
}

func (self *Writer) Tick() {
	self.intermittent = (self.intermittent + 1) % intermittentCycleTicks
	if self.ReachedEnd() { return }

	self.wait -= 1
	if self.wait <= 0 {
		char := self.text[self.index]
		if char >= 32 {
			switch char {
			case 44: self.wait += pauseStd*5 // ,
			case 46: self.wait += pauseStd*8 // .
			case 58: self.wait += pauseStd*8 // :
			case 33: self.wait += pauseStd*8 // !
			case 63: self.wait += pauseStd*8 // ?
			default:
				self.wait += waitStd
			}
		} else {
			switch char {
			case 3: self.wait += pauseStd*6
			case 4: self.wait += pauseStd*10
			case 5: self.wait += pauseStd*13
			default:
				self.wait = 0
			}
		}
		
		self.index += 1
		if self.wait != 0 {
			sound.PlaySFX(sound.SfxNav)
		}

		// quick fix to support utf8 for spanish and catalan
		for self.index < len(self.text) && self.text[self.index] >= 128 {
			self.index += 1
		}
	}
}

func (self *Writer) Draw(screen *ebiten.Image, fontSize int, bounds image.Rectangle, opacity uint8) (int, int) {
	self.renderer.SetTarget(screen)
	self.renderer.SetSizePx(fontSize)
	mainColor := color.RGBA{240, 240, 240, opacity}
	disabledColor := color.RGBA{140, 140, 140, opacity}
	nameColor := color.RGBA{255, 99, 112, opacity}
	self.renderer.SetColor(mainColor)

	shakeIntensity := 0
	intermittentOn := false
	comingFromLineBreak := true
	lineAdvance := self.renderer.GetLineAdvance()
	feed := self.renderer.NewFeed(fixed.P(bounds.Min.X, bounds.Min.Y))
	originalBreakX := feed.LineBreakX

	index := 0
	for index < self.index {
		ascii := self.text[index]
		if ascii < 32 {
			// special characters and control codes
			switch ascii {
			case 1: // use current x as next line start x
				feed.LineBreakX = feed.Position.X
			case 2: // reset next line start x to original value
				feed.LineBreakX = originalBreakX
			case 3, 4, 5: // 6 ticks pause
				// explicit pauses, already processed in Update()
			case 7:
				self.renderer.SetColor(mainColor)
			case 8:
				self.renderer.SetColor(disabledColor)
			case 9:
				self.renderer.SetColor(nameColor)
			case 10: // line break
				feed.LineBreak()
				comingFromLineBreak = true
			case 11: // 50% line break
				feed.Position.Y += lineAdvance/2
			case 12:
				shakeIntensity = 0
			case 13:
				shakeIntensity = 64
			case 14:
				shakeIntensity = 32
			case 16: // no-wait period. expected for robot names at start of line
				feed.Draw('.')
			case 17: // intermittent off
				intermittentOn = false
			case 18: // intermittent on
				intermittentOn = true
			default:
				panic(ascii)
			}
			index += 1
		} else if ascii == 32 {
			if !comingFromLineBreak { feed.Advance(' ') }
			index += 1
		} else { // draw next word
			comingFromLineBreak = false
			word := self.nextWord(index)

			// measure word to see if it fits
			width := self.renderer.SelectionRect(word).Width
			if (feed.Position.X + width).Ceil() > bounds.Max.X {
				feed.LineBreak() // didn't fit, jump to next line before drawing
			}

			// abort if we are going beyond the vertical working area
			if feed.Position.Y.Floor() >= bounds.Max.Y {
				return feed.Position.X.Ceil(), feed.Position.Y.Ceil()
			}

			// consider intermittency
			intrm := (intermittentOn && self.intermittent < intermittentCycleTicks/2)
			var clr color.RGBA
			if intrm {
				clr = self.renderer.GetColor().(color.RGBA)
				self.renderer.SetColor(color.RGBA{clr.R, clr.G, clr.B, clr.A - clr.A/4})
			}

			// draw the word character by character (so we can apply shaking)
			for charIndex, codePoint := range word {
				if index + charIndex >= self.index {
					return feed.Position.X.Ceil(), feed.Position.Y.Ceil()
				}
				if shakeIntensity > 0 {
					preY := feed.Position.Y
					vibr := fixed.Int26_6(rand.Intn(shakeIntensity))
					if rand.Intn(2) == 0 { vibr = -vibr }
					feed.Position.Y += vibr
					feed.Draw(codePoint)
					feed.Position.Y = preY
				} else {
					feed.Draw(codePoint)
				}
			}
			index += len(word)

			// restore color if intermittency active
			if intrm { self.renderer.SetColor(clr) }
		}

		// jump to next line if necessary
		if !comingFromLineBreak && feed.Position.X.Ceil() > bounds.Max.X {
			feed.LineBreak()
			comingFromLineBreak = true
		}
	}
	
	return feed.Position.X.Ceil(), feed.Position.Y.Ceil()
}

func (self *Writer) nextWord(index int) string {
	start := index
	for index < len(self.text) {
		if self.text[index] <= 32 { return self.text[start : index] }
		index += 1
	}
	return self.text[start : self.index]
}
