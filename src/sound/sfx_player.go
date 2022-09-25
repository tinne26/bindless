package sound

import "sync"
import "github.com/hajimehoshi/ebiten/v2/audio"

type SfxPlayer struct {
	mutex sync.Mutex
	context *audio.Context
	source []byte
	volume float64
}

func NewSfxPlayer(context *audio.Context, bytes []byte) *SfxPlayer{
	return &SfxPlayer { context: context, source: bytes, volume: 1.0 }
}

func (self *SfxPlayer) SetVolume(volume float64) {
	self.mutex.Lock()
	self.volume = volume
	self.mutex.Unlock()
}

func (self *SfxPlayer) Play() {
	player := self.context.NewPlayerFromBytes(self.source)
	self.mutex.Lock()
	player.SetVolume(self.volume)
	self.mutex.Unlock()
	player.Play()
}
