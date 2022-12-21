package main

// std library imports
import "os"
import "log"
import "math"
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

	// configure window
	// ebiten.SetWindowIcon(...)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Bindless")
	ebiten.SetScreenClearedEveryFrame(false)
	isFullscreen := true
	for _, arg := range os.Args {
		if arg == "--windowed" { isFullscreen = false }
		if arg == "--maxfps" {
			ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
			game.SetFPSDebugActive()
		}
	}
	ebiten.SetFullscreen(isFullscreen)

	// configure window size (you would have to keep track of device
	// scale factor and keep adjusting if you want to change device
	// scale factor through the game, but whatever...)
	// (actually, moving the window to another monitor can cause
	//  DPI scale to be changed, though... who wants to do that)
	scale := ebiten.DeviceScaleFactor()
	whole, fract := math.Modf(scale)
	if fract == 0 {
		ebiten.SetWindowSize(int(640*whole), int(360*whole))
	} else {
		cFactor := 1.0/(1.0 + fract)
		ebiten.SetWindowSize(int(640*whole*cFactor), int(360*whole*cFactor))
	}

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
