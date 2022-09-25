package episode

import "image"

import "github.com/hajimehoshi/ebiten/v2"

import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/misc/typewriter"
import "github.com/tinne26/bindless/src/art/palette"
import "github.com/tinne26/bindless/src/game/sceneitf"
import "github.com/tinne26/bindless/src/sound"

type fadeType uint8
const (
	fadeIn   fadeType = 0
	fadeNone fadeType = 1
	fadeOut  fadeType = 2
)

type episodeKey int
const (
	CleaningAutomaton episodeKey = 0
	ResearchLabDoor   episodeKey = 1
	ResearchLabGuard  episodeKey = 2
	ResearchLabSteal  episodeKey = 3
	JanaNewAbility    episodeKey = 4
	Infiltration      episodeKey = 5
	BasementDoor      episodeKey = 6
	InTheBasement     episodeKey = 7
)

type Episode struct {
	img *ebiten.Image
	writer *typewriter.Writer
	opacity uint8
	fade fadeType // fadeIn, fadeNone, fadeOut
	skipPressed bool
}

func New(ctx *misc.Context, key episodeKey) (*Episode, error) {
	ebiImg, err := misc.LoadPNG(ctx.Filesys, episodeImgPaths[int(key)])
	if err != nil { return nil, err }

	return &Episode {
		img: ebiImg,
		writer: typewriter.New(episodesRawText[int(key)].Get(), ctx),
	}, nil
}

func (self *Episode) Update(logCursorX, logCursorY int) error {

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

	switch self.fade {
	case fadeIn:
		if self.opacity >= 252 { self.opacity = 255 } else { self.opacity += 3 }
		if self.opacity == 255 { self.fade = fadeNone }
	case fadeNone:
		if self.writer.ReachedEnd() && misc.MousePressed() {
			self.fade = fadeOut
			sound.SfxClick.Play()
		}
	case fadeOut:
		if self.opacity <= 4 { self.opacity = 0 } else { self.opacity -= 4 }
	default:
		panic("invalid fade mode")
	}

	return nil
}

func (self *Episode) Draw(screen *ebiten.Image) {
	iw, ih := self.img.Size()
	sw, sh := screen.Size()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64((sw/2 - iw)/2), float64((sh - ih)/2))
	opts.ColorM.ScaleWithColor(palette.SceneWireframe)
	opts.ColorM.Scale(1.0, 1.0, 1.0, float64(self.opacity)/255)
	screen.DrawImage(self.img, opts)
}

func (self *Episode) DrawHiRes(screen *ebiten.Image, zoomLevel float64) {
	fontSize := misc.ScaledFontSize(10, zoomLevel)
	bounds := screen.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	horzMargin, vertMargin := w/28, h/10
	ox, oy := w/2 + horzMargin, vertMargin
	fx, fy := w - horzMargin/2, h - vertMargin/2
	rect := image.Rect(bounds.Min.X + ox, bounds.Min.Y + oy, bounds.Min.X + fx, bounds.Min.Y + fy)
	self.writer.Draw(screen, fontSize, rect, self.opacity)
}

func (self *Episode) Status() sceneitf.Status {
	if self.fade == fadeOut && self.opacity == 0 {
		return sceneitf.IsOverNext
	}
	return sceneitf.KeepAlive
}
