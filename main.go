package main

// std library imports
import "os"
import "log"
import "time"
import "embed"
import "math/rand"
import "runtime/pprof"

// external imports
import "github.com/hajimehoshi/ebiten/v2"

// internal imports
import "github.com/tinne26/bindless/src/misc"
import "github.com/tinne26/bindless/src/game"
import "github.com/tinne26/bindless/src/art/graphics"
import "github.com/tinne26/bindless/src/shaders"
import "github.com/tinne26/bindless/src/sound"

//go:embed assets/*
var assetsFS embed.FS

// seed rng
func init() { rand.Seed(time.Now().UnixNano()) }

const ProfileName = "profile.prof"
func main() {
	const profile = false

	misc.StartLoading()
	go func() {
		// preload some assets
		err := graphics.Load(&assetsFS)
		if err != nil { log.Fatal(err) }
		err = sound.Load(&assetsFS)
		if err != nil { log.Fatal(err) }

		// load shaders
		shaders.Load()

		// mark loading as complete
		misc.LoadingDone()
	}()

	// create game context (contains shared info like fontLib, filesys, etc.)
	ctx, err := misc.NewContext(&assetsFS)
	if err != nil { log.Fatal(err) }

	// create the main game handler
	game, err := game.New(ctx)
	if err != nil { log.Fatal(err) }

	// configure window and run the game
	// ebiten.SetWindowIcon(...)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Bindless")
	ebiten.SetWindowSize(640, 360)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetFullscreen(true)

	var file *os.File
	if profile {
		// prepare profiling file
		file, err = os.Create(ProfileName)
		if err != nil { log.Fatal(err.Error()) }
		ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
		ebiten.SetFullscreen(true)

		// start profiling
		err = pprof.StartCPUProfile(file)
		if err != nil { log.Fatal(err.Error()) }
	}

	err = ebiten.RunGame(game)
	if err != nil { log.Fatal(err) }

	if profile {
		// stop profiling
		pprof.StopCPUProfile()
		file.Close()
	}
}
