package sound

//import "log"
import "io"
import "sync"
import "bytes"

import "github.com/hajimehoshi/ebiten/v2/audio/mp3"

type Loop struct {
	StartPosition int64
	LoopBackId uint16
	RepeatsLeft uint16
}

type Looper struct {
	stream io.ReadSeeker
	mutex sync.Mutex
	position int64
	loopStart int64
	activeLoopEnd int64
	loopEnd int64
}

func NewLooper(stream io.ReadSeeker, loopStart int64, loopEnd int64) *Looper {
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
	case io.SeekEnd:
		switch stream := self.stream.(type) {
		case *bytes.Reader:
			startOffset = int64(stream.Len()) + offset
		case *mp3.Stream:
			startOffset = stream.Length() + offset
		// case *ogg.Stream:
		// 	startOffset = stream.Length() + offset
		default:
			startOffset = self.activeLoopEnd + offset
		}
	}

	whence = io.SeekStart
	if startOffset > self.activeLoopEnd {
		offset = self.loopStart
	} else if startOffset < 0 {
		offset = 0
	}

	var err error
	self.position, err = self.stream.Seek(offset, whence)
	return self.position, err
}

func (self *Looper) Reset(stream io.ReadSeeker, loopStart int64, loopEnd int64) {
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
