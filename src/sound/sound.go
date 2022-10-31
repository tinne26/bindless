package sound

import "io"
import "io/ioutil"
import "embed"

import "github.com/hajimehoshi/ebiten/v2/audio"
import "github.com/hajimehoshi/ebiten/v2/audio/vorbis"

var ObsessiveMechanics io.ReadSeeker
var MagneticCityMemories io.ReadSeeker
var MeddlesomeTheory io.ReadSeeker
var obsessiveShortLoop bool = false

// to play sfx, use sound.SfxNav.Play() or similar
var SfxTypeA *SfxPlayer
var SfxTypeB *SfxPlayer
var SfxTypeC *SfxPlayer
var SfxNav *SfxPlayer
var SfxTileNav *SfxPlayer
var SfxAbility *SfxPlayer
var SfxNope *SfxPlayer
var SfxClick *SfxPlayer

var bgmMaxVol float64 = 0.5

var bgmVolume float64 = 0
var bgmFadeTarget float64
var bgmFadeSpeed float64 = 0.002
var bgmNextStream io.ReadSeeker

var ctx *audio.Context
var bgmPlayer *audio.Player
var bgmLooper *Looper
var activeStream io.ReadSeeker

func Load(filesys *embed.FS) error {
	ctx = audio.NewContext(44100)

	folder := "assets/audio/"
	file, err := filesys.Open(folder + "obsessive_mechanics.ogg")
	if err != nil { return err }
	ObsessiveMechanics, err = vorbis.DecodeWithSampleRate(44100, file)
	if err != nil { return err }

	file, err = filesys.Open(folder + "magnetic_city_memories.ogg")
	if err != nil { return err }
	MagneticCityMemories, err = vorbis.DecodeWithSampleRate(44100, file)
	if err != nil { return err }

	file, err = filesys.Open(folder + "meddlesome_theory.ogg")
	if err != nil { return err }
	MeddlesomeTheory, err = vorbis.DecodeWithSampleRate(44100, file)
	if err != nil { return err }

	// load sfx
	folder += "sfx/"
	sfxBytes, err := loadAudioBytes(filesys, folder + "tile_nav.ogg")
	if err != nil { return err }
	SfxNav = NewSfxPlayer(ctx, sfxBytes)
	SfxNav.SetVolume(0.4)
	SfxTileNav = NewSfxPlayer(ctx, sfxBytes)
	SfxTileNav.SetVolume(0.36)

	sfxBytes, err = loadAudioBytes(filesys, folder + "nope.ogg")
	if err != nil { return err }
	SfxNope = NewSfxPlayer(ctx, sfxBytes)

	sfxBytes, err = loadAudioBytes(filesys, folder + "ability.ogg")
	if err != nil { return err }
	SfxAbility = NewSfxPlayer(ctx, sfxBytes)

	sfxBytes, err = loadAudioBytes(filesys, folder + "click.ogg")
	if err != nil { return err }
	SfxClick = NewSfxPlayer(ctx, sfxBytes)
	SfxClick.SetVolume(0.74)

	sfxBytes, err = loadAudioBytes(filesys, folder + "type_a.ogg")
	if err != nil { return err }
	SfxTypeA = NewSfxPlayer(ctx, sfxBytes)
	sfxBytes, err = loadAudioBytes(filesys, folder + "type_b.ogg")
	if err != nil { return err }
	SfxTypeB = NewSfxPlayer(ctx, sfxBytes)
	sfxBytes, err = loadAudioBytes(filesys, folder + "type_c.ogg")
	if err != nil { return err }
	SfxTypeC = NewSfxPlayer(ctx, sfxBytes)
	SfxTypeA.SetVolume(0.34)
	SfxTypeB.SetVolume(0.34)
	SfxTypeC.SetVolume(0.34)

	return nil
}

func Update() {
	// bgm volume transitions
	if bgmVolume < bgmFadeTarget {
		bgmVolume += bgmFadeSpeed
		if bgmVolume > bgmFadeTarget {
			bgmVolume = bgmFadeTarget
		}
		bgmPlayer.SetVolume(bgmVolume)
	} else if bgmVolume > bgmFadeTarget {
		bgmVolume -= bgmFadeSpeed
		if bgmVolume < 0 { bgmVolume = 0 }

		if bgmNextStream != nil {
			setupNextStream()
		} else if bgmPlayer != nil {
			bgmPlayer.SetVolume(bgmVolume)
		}
	} else if bgmVolume == 0 && bgmNextStream != nil {
		setupNextStream()
	}
}

func setupNextStream() {
	activeStream = bgmNextStream
	needsReset := true
	var volFactor float64
	switch activeStream {
	case MagneticCityMemories:
		setLoopPoints(0, 6379452*4, needsReset)
		volFactor = 1.0
	case ObsessiveMechanics:
		_, loopEnd := obsessiveLoopPoints()
		setLoopPoints(0, loopEnd, needsReset)
		volFactor = 1.17
	case MeddlesomeTheory:
		setLoopPoints(7456*4, 2295900*4, needsReset)
		volFactor = 0.8
	default:
		panic("unexpected stream")
	}

	if bgmPlayer != nil {
		bgmPlayer.Pause()
		err := bgmPlayer.Close()
		if err != nil { panic(err) }
	}

	var err error
	bgmPlayer, err = ctx.NewPlayer(bgmLooper)
	if err != nil { panic(err) }

	bgmNextStream = nil
	bgmPlayer.SetVolume(0)
	bgmPlayer.Play()
	bgmFadeTarget = bgmMaxVol*volFactor
}

func setLoopPoints(loopStart, loopEnd int64, needsReset bool) {
	if bgmLooper == nil {
		bgmLooper = NewLooper(activeStream, loopStart, loopEnd)
	} else if needsReset {
		bgmLooper.Reset(activeStream, loopStart, loopEnd)
	} else {
		bgmLooper.AdjustLoop(loopStart, loopEnd)
	}
}

func RequestBGM(stream io.ReadSeeker) {
	if activeStream == stream && bgmFadeTarget != 0 { return }
	bgmFadeTarget = 0
	bgmNextStream = stream
}

func SetBGMFadeSpeed(speed float64) {
	bgmFadeSpeed = speed // concurrent access, you were saying?
}

func RequestFadeOut() {
	bgmFadeTarget = 0
}

func SetObssessiveShortLoop(active bool) {
	if obsessiveShortLoop != active {
		obsessiveShortLoop = active
		if activeStream == ObsessiveMechanics {
			loopStart, loopEnd := obsessiveLoopPoints()
			needsReset := false
			setLoopPoints(loopStart, loopEnd, needsReset)
		}
	}
}

func obsessiveLoopPoints() (int64, int64) {
	if obsessiveShortLoop {
		return 0, 1283320*4
	} else {
		return 1283337*4, 4491300*4
	}
}

func loadAudioBytes(filesys *embed.FS, filename string) ([]byte, error) {
	file, err := filesys.Open(filename)
	if err != nil { return nil, err }
	stream, err := vorbis.DecodeWithSampleRate(44100, file)
	if err != nil { return nil, err }
	audioBytes, err := ioutil.ReadAll(stream)
	return audioBytes, err
}
