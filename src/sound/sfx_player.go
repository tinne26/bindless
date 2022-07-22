package sound

import "github.com/hajimehoshi/ebiten/v2/audio"

type SfxPlayer struct {
	context *audio.Context
	source []byte
}

func NewSfxPlayer(context *audio.Context, bytes []byte) *SfxPlayer{
	return &SfxPlayer { context: context, source: bytes }
}

func (self *SfxPlayer) NewPlayWithVolume(volume float64) {
	player := self.context.NewPlayerFromBytes(self.source)
	player.SetVolume(volume)
	player.Play()
}
