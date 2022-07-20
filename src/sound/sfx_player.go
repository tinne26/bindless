package sound

import "sync"

import "github.com/hajimehoshi/ebiten/v2/audio"

type SfxPlayer struct {
	players sync.Pool
	activePlayers []*audio.Player
}

func NewSfxPlayer(context *audio.Context, bytes []byte) *SfxPlayer{
	pool := sync.Pool {
		New: func() any {
			return context.NewPlayerFromBytes(bytes)
		}}
	return &SfxPlayer { players: pool, activePlayers: make([]*audio.Player, 0, 1) }
}

func (self *SfxPlayer) NewPlayWithVolume(volume float64) {
	player := self.players.Get().(*audio.Player)
	player.SetVolume(volume)
	player.Play()
	self.activePlayers = append(self.activePlayers, player)
}

func (self *SfxPlayer) Update() {
	i := 0
	for i < len(self.activePlayers) {
		player := self.activePlayers[i]
		if !player.IsPlaying() {
			err := player.Rewind()
			if err != nil { panic(err) }
			self.players.Put(player)
			self.removeActivePlayerAt(i)
		} else {
			i += 1
		}
	}
}

func (self *SfxPlayer) removeActivePlayerAt(i int) {
	playerCount := len(self.activePlayers)
	if i >= playerCount { panic("index out of bounds") }
	if playerCount == 1 {
		self.activePlayers = self.activePlayers[0 : 0]
	} else {
		self.activePlayers[i], self.activePlayers[playerCount - 1] = self.activePlayers[playerCount - 1], self.activePlayers[i]
		self.activePlayers = self.activePlayers[:playerCount - 1]
	}
}
