package main

// std library imports
import "log"
import "time"
import "embed"
import "math/rand"

// external imports
import "github.com/hajimehoshi/ebiten/v2"

// internal imports
import "bindless/src/misc"
import "bindless/src/game"
import "bindless/src/art/graphics"
import "bindless/src/shaders"
import "bindless/src/sound"

//go:embed assets/*
var assetsFS embed.FS

// seed rng
func init() { rand.Seed(time.Now().UnixNano()) }

func main() {
	// preload some assets
	err := graphics.Load(&assetsFS)
	if err != nil { log.Fatal(err) }
	err = sound.Load(&assetsFS)
	if err != nil { log.Fatal(err) }

	// create game context (contains shared info like fontLib, filesys, etc.)
	ctx, err := misc.NewContext(&assetsFS)
	if err != nil { log.Fatal(err) }

	// load shaders
	shaders.Load()

	// create the main game handler
	game, err := game.New(ctx)
	if err != nil { log.Fatal(err) }

	// configure window and run the game
	// ebiten.SetWindowIcon(...)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	//ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowTitle("Bindless")
	ebiten.SetWindowSize(640, 360)
	err = ebiten.RunGame(game)
	if err != nil { log.Fatal(err) }
}
