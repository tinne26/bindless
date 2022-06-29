package game

// std library imports
import "io"
import "math"
import "image"
import "strconv"
import "image/color"

// external imports
import "github.com/hajimehoshi/ebiten/v2"
import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

// internal imports
import "bindless/src/misc"
import "bindless/src/misc/background"
import "bindless/src/game/sceneitf"
import "bindless/src/scenes/title"
import "bindless/src/scenes/text"
import "bindless/src/scenes/episode"
import "bindless/src/scenes/level"
import "bindless/src/sound"

const numScenes = 21

// *Game implements the ebiten.Game interface
type Game struct {
	context *misc.Context
	logicalScreen *ebiten.Image

	background *background.Background
	scene sceneitf.Scene
	sceneId int

	quickLevelJumpIndex int
	lastJumpKeyPress ebiten.Key

	disableHighDPI bool
	fKeyPressed bool
	dKeyPressed bool
	fpsDebugActive bool

	prevHiResRect image.Rectangle // used to map mouse coords from hiRes to logical coords
}

func New(ctx *misc.Context) (*Game, error) {
	game := &Game{
		context: ctx,
		logicalScreen: ebiten.NewImage(640, 360),
		scene: nil,
	}

	return game, nil
}

func (self *Game) Layout(w, h int) (int, int) {
	if self.disableHighDPI { return w, h }
	factor := ebiten.DeviceScaleFactor()
	return int(float64(w)*factor), int(float64(h)*factor)
}

func (self *Game) Update() error {
	// first scene load when asset loading is done
	if !misc.IsLoadingDone() { return nil }
	if self.background == nil {
		self.background = background.New()
		return self.loadScene(0)
	}

	// sound update
	sound.Update()

	// scene jump hack
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		for n, key := range []ebiten.Key{ ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9 } {
			if ebiten.IsKeyPressed(key) {
				if self.lastJumpKeyPress != key {
					self.lastJumpKeyPress = key
					self.quickLevelJumpIndex *= 10
					self.quickLevelJumpIndex += n
				}
			} else if self.lastJumpKeyPress == key {
				self.lastJumpKeyPress = ebiten.KeyEscape
			}
		}
	} else {
		self.lastJumpKeyPress = ebiten.KeyEscape
		if self.quickLevelJumpIndex != 0 {
			sceneId := self.quickLevelJumpIndex - 1
			self.quickLevelJumpIndex = 0
			if sceneId < numScenes {
				err := self.loadScene(sceneId)
				return err
			}
		}
	}

	// fullscreen update
	fKeyPressed := ebiten.IsKeyPressed(ebiten.KeyF)
	if !fKeyPressed {
		self.fKeyPressed = false
	} else if !self.fKeyPressed {
		self.fKeyPressed = true
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	// fps debug
	dKeyPressed := ebiten.IsKeyPressed(ebiten.KeyD)
	if !dKeyPressed {
		self.dKeyPressed = false
	} else if !self.dKeyPressed {
		self.dKeyPressed = true
		self.fpsDebugActive = !self.fpsDebugActive
	}

	// update background
	self.background.Update()

	// handle scene status
	switch self.scene.Status() {
	case sceneitf.KeepAlive:
		// ok, nothing to do
	case sceneitf.IsOver: // go to next scene
		err := self.loadScene((self.sceneId + 1) % numScenes)
		if err != nil { return err }
	case sceneitf.Restart: // start scene again
		err := self.loadScene(self.sceneId)
		if err != nil { return err }
	}

	// update scene
	cx, cy := ebiten.CursorPosition()
	prevWidth, prevHeight := float64(self.prevHiResRect.Dx()), float64(self.prevHiResRect.Dy())
	logCursorX := int(float64(cx - self.prevHiResRect.Min.X)*(640.0/prevWidth))
	logCursorY := int(float64(cy - self.prevHiResRect.Min.Y)*(360.0/prevHeight))
	return self.scene.Update(logCursorX, logCursorY)
}

func (self *Game) Draw(screen *ebiten.Image) {
	if self.background == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "loading game assets, please wait...")
		return
	}

	// draw background and scene to the logical screen
	//self.logicalScreen.Clear() // no need, background already draws this
	self.background.Draw(self.logicalScreen)
	self.scene.Draw(self.logicalScreen)

	// draw logical screen onto the main screen, scaling as needed.
	// there's quite a lot of zoomLevel calculations and stuff
	w, h := screen.Size()
	xZoom, yZoom := float64(w)/640.0, float64(h)/360.0
	zoomLevel := math.Min(xZoom, yZoom)
	intZoomLevel := int(zoomLevel)
	propWidth, propHeight := 640*intZoomLevel, 360*intZoomLevel

	// ...determine margins
	marginHorz := (w - propWidth)/2
	marginVert := (h - propHeight)/2

	// ...configure draw image options and issue DrawImage()
	opts := &ebiten.DrawImageOptions{}
	xScale, yScale := float64(intZoomLevel), float64(intZoomLevel)
	if zoomLevel < 1 {
		// game doesn't fit in its expected minimum size,
		// downscale even if this will be almost unplayable
		marginHorz, marginVert = 0, 0
		xScale, yScale = float64(w)/640.0, float64(h)/360.0
		if w >= 640 { xScale = 1.0; marginHorz = (w - 640)/2 }
		if h >= 360 { yScale = 1.0; marginVert = (h - 360)/2 }
		opts.Filter = ebiten.FilterLinear
	}
	opts.GeoM.Scale(xScale, yScale)
	if marginHorz != 0 || marginVert != 0 { // fill margins if needed
		screen.Fill(color.RGBA{24, 24, 24, 255})
		opts.GeoM.Translate(float64(marginHorz), float64(marginVert))
	}
	screen.DrawImage(self.logicalScreen, opts) // finally

	// draw high resolution scene elements
	if zoomLevel >= 1 {
		self.prevHiResRect = image.Rect(marginHorz, marginVert, propWidth + marginHorz, propHeight + marginVert)
		zoomLevel = float64(intZoomLevel)
	} else {
		uw, uh := int(xScale*640), int(yScale*360)
		self.prevHiResRect = image.Rect(marginHorz, marginVert, uw + marginHorz, uh + marginVert)
	}
	hiResScreen := screen.SubImage(self.prevHiResRect).(*ebiten.Image)
	self.scene.DrawHiRes(hiResScreen, zoomLevel)

	// fps debug
	if self.fpsDebugActive {
		ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.CurrentFPS(), 'f', 2, 64) + "fps")
	}
}

func (self *Game) loadScene(id int) error {
	var err error
	switch id {
	case 0: // title screen
		sound.SetBGMFadeSpeed(0.001)
		sound.RequestBGM(sound.MagneticCityMemories)
		self.scene, err = title.New(self.context)
		if err != nil { return err }
	case 1: // preamble
		cfgLevelSound(sound.MagneticCityMemories)
		sound.SetBGMFadeSpeed(0.001)
		self.scene, err = text.New(self.context, text.Preamble)
		if err != nil { return err }
	case 2: // first scene (cleaning automaton)
		cfgLevelSound(sound.ObsessiveMechanics)
		sound.SetObssessiveShortLoop(true)
		self.scene, err = episode.New(self.context, episode.CleaningAutomatonTest)
		if err != nil { return err }
	case 3: // first level, tutorial pt1
		cfgLevelSound(sound.ObsessiveMechanics)
		self.scene, err = level.New(self.context, level.CleanerTestDock)
		if err != nil { return err }
	case 4: // tutorial pt2
		cfgLevelSound(sound.ObsessiveMechanics)
		self.scene, err = level.New(self.context, level.CleanerTestRewire)
		if err != nil { return err }
	case 5: // first real level
		sound.RequestBGM(sound.ObsessiveMechanics)
		self.scene, err = level.New(self.context, level.CleanerTestReal)
		if err != nil { return err }
	case 6: // episode research lab door
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = episode.New(self.context, episode.ResearchLabDoor)
		if err != nil { return err }
	case 7: // research lab door
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = level.New(self.context, level.ResearchLabDoor)
		if err != nil { return err }
	case 8: // episode research lab guard
		cfgLevelSound(sound.ObsessiveMechanics)
		sound.SetObssessiveShortLoop(true)
		self.scene, err = episode.New(self.context, episode.ResearchLabGuard)
		if err != nil { return err }
	case 9: // first guard layer
		cfgLevelSound(sound.ObsessiveMechanics)
		self.scene, err = level.New(self.context, level.ResearchLabGuard1)
		if err != nil { return err }
	case 10: // second guard layer
		cfgLevelSound(sound.ObsessiveMechanics)
		self.scene, err = level.New(self.context, level.ResearchLabGuard2)
		if err != nil { return err }
	case 11: // episode research lab
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = episode.New(self.context, episode.ResearchLabSteal)
		if err != nil { return err }
	case 12: // episode jana's note and modified MSP
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = episode.New(self.context, episode.JanaNewAbility)
		if err != nil { return err }
	case 13: // switch test
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = level.New(self.context, level.SwitchTest)
		if err != nil { return err }
	case 14: // episode (failed) infiltration
		cfgLevelSound(sound.ObsessiveMechanics)
		sound.SetObssessiveShortLoop(true)
		self.scene, err = episode.New(self.context, episode.Infiltration)
		if err != nil { return err }
	case 15: // infiltration guard level (this one is so cool)
		cfgLevelSound(sound.ObsessiveMechanics)
		self.scene, err = level.New(self.context, level.FinalGuard)
		if err != nil { return err }
	case 16: // episode basement door
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = episode.New(self.context, episode.BasementDoor)
		if err != nil { return err }
	case 17: // final door level
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = level.New(self.context, level.FinalDoor)
		if err != nil { return err }
	case 18: // episode in the basement
		cfgLevelSound(sound.MagneticCityMemories)
		self.scene, err = episode.New(self.context, episode.InTheBasement)
		if err != nil { return err }
	case 19: // to be continued
		cfgLevelSound(sound.MagneticCityMemories)
		sound.SetBGMFadeSpeed(0.001)
		self.scene, err = text.New(self.context, text.ToBeContinued)
		if err != nil { return err }
	case 20: // final words
		sound.SetBGMFadeSpeed(0.001)
		sound.RequestFadeOut()
		self.scene, err = text.New(self.context, text.Afterword)
		if err != nil { return err }
	default:
		panic("invalid scene id")
	}
	self.sceneId = id
	return nil
}

func cfgLevelSound(stream io.ReadSeeker) {
	sound.SetBGMFadeSpeed(0.05) // use fast sound transitions
	sound.RequestBGM(stream)
	if stream == sound.ObsessiveMechanics {
		sound.SetObssessiveShortLoop(false)
	}
}
