package background

// std library imports
import "math/rand"

// external imports
import "github.com/hajimehoshi/ebiten/v2"

// internal imports
import "bindless/src/art/graphics"
import "bindless/src/art/palette"

type backDecoration struct {
	x int16
	y int16
	stayTicks uint16
	opacity float64 // from 0 to 1
	fadeChange float64
	imgType uint8
}

func newBackDecoration() backDecoration {
	decoration := backDecoration{}
	decoration.Restart()
	return decoration
}

func (self *backDecoration) Restart() {
	self.x = int16(rand.Intn(660) - 20)
	self.y = int16(rand.Intn(380) - 20)
	self.imgType = uint8(rand.Intn(9))
	self.opacity = 0
	self.stayTicks = uint16(rand.Intn(120) + 80)
	self.fadeChange = rand.Float64()*0.02 + 0.004
}

func (self *backDecoration) Update() {
	// fading in case
	if self.fadeChange > 0 {
		self.opacity += self.fadeChange
		if self.opacity > 1 {
			self.opacity = 1
			self.fadeChange = -self.fadeChange
		}
		return
	}

	// stay case
	if self.stayTicks > 0 {
		self.stayTicks -= 1
		return
	}

	// fade out case
	self.opacity += self.fadeChange
	if self.opacity <= 0 {
		self.Restart()
	}
}

func (self *backDecoration) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.ColorM.Scale(1, 1, 1, self.opacity)
	opts.GeoM.Translate(float64(self.x), float64(self.y))
	screen.DrawImage(graphics.BackDecorations[self.imgType], opts)
}

// the background in the game is a plain color and many small
// decorative geometric images fading in and out
type Background struct {
	decorations []backDecoration
}

func New() *Background {
	const NumDecors = 54

	decors := make([]backDecoration, 0, NumDecors)
	for i := 0; i < NumDecors; i++ { decors = append(decors, newBackDecoration()) }
	return &Background { decorations: decors }
}

func (self *Background) Update() {
	for i := 0; i < len(self.decorations); i++ {
		self.decorations[i].Update()
	}
}

func (self *Background) Draw(screen *ebiten.Image) {
	screen.Fill(palette.Background)
	for _, decor := range self.decorations {
		decor.Draw(screen)
	}
}
