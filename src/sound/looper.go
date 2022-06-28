package sound
//import "log"
import "io"
import "sync"

import "github.com/hajimehoshi/ebiten/v2/audio/mp3"

// an audio looper has a base structure, and then a copy that's instanced
// from that structure, and when it ends it resets itself.
// the read implementation is:
// while bytesLeft > 0
// 	bytes, err := readCurrentPart(maxBytes)
//    if err != nil { return err }
//    bytesLeft -= len(bytes)
// MultiLoop
// struct:
// > stream *mp3.Stream
// > position int
// > structure []loops
// > parts []loops
// > currentPart int
//    >> where loops have: startPosition int, loopBackId int, repeatsLeft int
// Methods, besides Read and Seek... SetStructure(), etc.

type Loop struct {
	StartPosition int64
	LoopBackId uint16
	RepeatsLeft uint16
}

type Looper struct {
	stream *mp3.Stream
	mutex sync.Mutex
	position int64
	loopStart int64
	activeLoopEnd int64
	loopEnd int64
}

func NewLooper(stream *mp3.Stream, loopStart int64, loopEnd int64) *Looper {
	return &Looper {
		stream: stream,
		loopStart: loopStart,
		loopEnd: loopEnd,
		activeLoopEnd: loopEnd,
	}
}

func (self *Looper) Read(buffer []byte) (int, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	untilNextLoop := self.activeLoopEnd - self.position
	if untilNextLoop < 0 {
		//log.Printf("loop skip")
		untilNextLoop = 0
	}
	if int64(len(buffer)) > untilNextLoop {
		var err error
		var n int
		if untilNextLoop > 0 {
			n, err = self.stream.Read(buffer[0 : untilNextLoop])
			self.position += int64(n)
			if err != nil { return n, err }
		}
		self.activeLoopEnd = self.loopEnd
		self.position, err = self.stream.Seek(self.loopStart, io.SeekStart)
		if err != nil { return n, err }

		n, err = self.stream.Read(buffer[untilNextLoop : len(buffer)])
		self.position += int64(n)
		return n, err
	} else {
		n, err := self.stream.Read(buffer)
		self.position += int64(n)
		return n, err
	}
}

func (self *Looper) ChangeLoop(loopStart int64, loopEnd int64) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.loopStart = loopStart
	self.loopEnd = loopEnd
	if loopEnd >= self.position {
		self.activeLoopEnd = loopEnd
	}
}

func (self *Looper) Seek(offset int64, whence int) (int64, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	pos, err := self.lockedSeek(offset, whence)
	return pos, err
}

func (self *Looper) lockedSeek(offset int64, whence int) (int64, error) {
	startOffset := int64(0)
	switch whence {
	case io.SeekStart:   startOffset = offset
	case io.SeekCurrent: startOffset = self.position + offset
	case io.SeekEnd:     startOffset = self.stream.Length() + offset
	}
	if startOffset > self.activeLoopEnd {
		//return 0, fmt.Errorf("can't seek beyond loopEnd")
		//log.Printf("loop skip with whence = %d", whence)
	}

	var err error
	self.position, err = self.stream.Seek(offset, whence)
	return self.position, err
}

func (self *Looper) Reset(stream *mp3.Stream, loopStart int64, loopEnd int64) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	self.stream = stream
	self.loopStart = loopStart
	self.loopEnd = loopEnd
	self.activeLoopEnd = loopEnd
	self.position = 0

	pos, err := self.lockedSeek(0, io.SeekStart)
	if err != nil { panic(err) }
	if pos != 0 { panic("non-zero position after stream reset") }
}
