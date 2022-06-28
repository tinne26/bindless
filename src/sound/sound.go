package sound
// import "log"
// import "time"
import "embed"
import "math/rand"

import "github.com/hajimehoshi/ebiten/v2/audio"
import "github.com/hajimehoshi/ebiten/v2/audio/mp3"

var ObsessiveMechanics *mp3.Stream
var MagneticCityMemories *mp3.Stream
var obsessiveShortLoop bool = false

// to play sfx, use sound.PlaySfx(sound.SfxNav)
var SfxNav *audio.Player
var SfxLoudNav *audio.Player
var SfxAbility *audio.Player
var SfxNope *audio.Player
var SfxClick *audio.Player

var bgmMaxVol float64 = 0.7
var bgmVolume float64 = 0
var bgmFadeTarget float64
var bgmFadeSpeed float64 = 0.002
var bgmNextStream *mp3.Stream

var ctx *audio.Context
var bgmPlayer *audio.Player
var bgmLooper *Looper
var activeStream *mp3.Stream

func Load(filesys *embed.FS) error {
	ctx = audio.NewContext(44100)
	bgmLooper = NewLooper(nil, 0, 0)

	folder := "assets/audio/"
	file, err := filesys.Open(folder + "obsessive_mechanics.mp3")
	if err != nil { return err }
	ObsessiveMechanics, err = mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }

	file, err = filesys.Open(folder + "magnetic_city_memories.mp3")
	if err != nil { return err }
	MagneticCityMemories, err = mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }

	// load sfx
	folder += "sfx/"
	file, err = filesys.Open(folder + "nav.mp3")
	if err != nil { return err }
	sfx, err := mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }
	SfxNav, err = ctx.NewPlayer(sfx)
	if err != nil { return err }

	sfx, err = mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }
	SfxLoudNav, err = ctx.NewPlayer(sfx)
	if err != nil { return err }
	SfxLoudNav.SetVolume(0.5)

	file, err = filesys.Open(folder + "nope.mp3")
	if err != nil { return err }
	sfx, err = mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }
	SfxNope, err = ctx.NewPlayer(sfx)
	if err != nil { return err }

	file, err = filesys.Open(folder + "ability.mp3")
	if err != nil { return err }
	sfx, err = mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }
	SfxAbility, err = ctx.NewPlayer(sfx)
	if err != nil { return err }

	file, err = filesys.Open(folder + "click.mp3")
	if err != nil { return err }
	sfx, err = mp3.DecodeWithSampleRate(44100, file)
	if err != nil { return err }
	SfxClick, err = ctx.NewPlayer(sfx)
	if err != nil { return err }

	return nil
}

func Update() {
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
	var err error

	// the time loop calculations don't make any sense at all. it's ok,
	// don't try to understand it, there's a bug somewhere
	switch activeStream {
	case MagneticCityMemories:
		bgmLooper.Reset(activeStream, 0, msToByte(144659 + 500))
	case ObsessiveMechanics:
		bgmLooper.Reset(activeStream, 0, msToByte(101814 + 500))
		cfgObsessiveLoop()
	default:
		panic("unexpected stream")
	}

	if bgmPlayer != nil {
		bgmPlayer.Pause()
		err := bgmPlayer.Close()
		if err != nil { panic(err) }
	}
	bgmPlayer, err = ctx.NewPlayer(bgmLooper)
	if err != nil { panic(err) }

	// samples := (activeStream.Length()) / 4
	// targetTime := time.Duration(samples)*time.Second/time.Duration(44100)
	// targetTime -= time.Second*4
	//
	// err = bgmPlayer.Seek(targetTime) // 25*time.Second)
	// if err != nil { log.Fatal(err) }

	bgmNextStream = nil
	bgmPlayer.SetVolume(0)
	bgmPlayer.Play()
	bgmFadeTarget = bgmMaxVol
	if activeStream == ObsessiveMechanics {
		bgmFadeTarget -= 0.24 // fix for loudness
	}
}

func msToByte(ms int) int64 {
	nearest := int64(176.4*float64(ms))
	fract := nearest % 4
	return nearest - fract
}

func RequestBGM(stream *mp3.Stream) {
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

func PlaySFX(sfxPlayer *audio.Player) {
	err := sfxPlayer.Rewind()
	if err != nil { panic(err) }
	if sfxPlayer == SfxNav {
		sfxPlayer.SetVolume(0.34 + rand.Float64()/16.0)
	}
	sfxPlayer.Play()
}

func SetObssessiveShortLoop(active bool) {
	if obsessiveShortLoop != active {
		obsessiveShortLoop = active
		cfgObsessiveLoop()
	}
}

func cfgObsessiveLoop() {
	if activeStream != ObsessiveMechanics { return }
	if obsessiveShortLoop {
		bgmLooper.ChangeLoop(0, msToByte(29091 + 500))
	} else {
		bgmLooper.ChangeLoop(msToByte(29091), msToByte(101814 + 500))
	}
}
