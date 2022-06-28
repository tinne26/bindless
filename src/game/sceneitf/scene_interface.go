package sceneitf

import "github.com/hajimehoshi/ebiten/v2"

type Status uint8
const (
	KeepAlive Status = 0
	IsOver    Status = 1
	Restart   Status = 2
)

// A simple interface that all scenes in bindless/src/scene implicitly
// implement. This allows the main Game interface to operate with them
// and switch scenes more easily.
type Scene interface {
	Update(logCursorX, logCursorY int) error
	Draw(screen *ebiten.Image)
	DrawHiRes(screen *ebiten.Image, zoomLevel float64)
	Status() Status
}
