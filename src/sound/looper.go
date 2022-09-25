package sound

import "io"
import "sync"
import "bytes"

// implementation taken from tinne26/edau.Looper

type Looper struct {
	stream io.ReadSeeker
	mutex sync.Mutex
	position int64
	loopStart int64
	activeLoopEnd int64 // relevant when modifying loop points. if we are going towards an
	                    // end loop point but we change it for another that comes earlier
							  // but we have already past it, we still have to continue towards
							  // the previous loop end point
	loopEnd int64
}

func NewLooper(stream io.ReadSeeker, loopStart int64, loopEnd int64) *Looper {
	assertLoopValuesValidity(loopStart, loopEnd)
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

	var bytesRead int
	for len(buffer) > 0 {
		untilNextLoop := self.activeLoopEnd - self.position
		if untilNextLoop < 0 { untilNextLoop = 0 }
		
		// simple case: not reaching next loop point yet
		if int64(len(buffer)) <= untilNextLoop {
			n, err := self.readAll(buffer)
			bytesRead += n
			return bytesRead, err
		}
	
		// complex case: one or more loop points reached
		if untilNextLoop > 0 {
			n, err := self.readAll(buffer[0 : untilNextLoop])
			bytesRead += n
			if err != nil { return bytesRead, err }
		}
	
		var err error
		self.activeLoopEnd = self.loopEnd
		self.position, err = self.stream.Seek(self.loopStart, io.SeekStart)
		if err != nil { return bytesRead, err }
		buffer = buffer[untilNextLoop : ]
	}

	return bytesRead, nil
}

func (self *Looper) readAll(buffer []byte) (int, error) {
	bytesRead := 0
	defer func(){ self.position += int64(bytesRead) }()
	
	for {
		// read and return if we are done or got an error
		n, err := self.stream.Read(buffer)
		bytesRead += n
		if n == len(buffer) || err != nil {
			return bytesRead, err
		}

		// we didn't read enough, try again
		buffer = buffer[n : ]
	}
}

func (self *Looper) Seek(offset int64, whence int) (int64, error) {
	self.mutex.Lock()
	n, err := self.stream.Seek(offset, whence)
	self.position = n
	if self.position <= self.loopEnd {
		self.activeLoopEnd = self.loopEnd
	}
	self.mutex.Unlock()
	return n, err
}

func (self *Looper) GetPosition() int64 {
	self.mutex.Lock()
	position := self.position
	self.mutex.Unlock()
	return position
}

func (self *Looper) GetLoopStart() int64 {
	self.mutex.Lock()
	loopStart := self.loopStart
	self.mutex.Unlock()
	return loopStart
}

func (self *Looper) GetLoopEnd() int64 {
	self.mutex.Lock()
	loopEnd := self.loopEnd
	defer self.mutex.Unlock()
	return loopEnd
}

func (self *Looper) GetLoopPoints() (int64, int64) {
	self.mutex.Lock()
	loopStart := self.loopStart
	loopEnd   := self.loopEnd
	self.mutex.Unlock()
	return loopStart, loopEnd
}

func (self *Looper) AdjustLoop(loopStart, loopEnd int64) {
	assertLoopValuesValidity(loopStart, loopEnd)
	self.mutex.Lock()
	self.loopStart = loopStart
	self.loopEnd = loopEnd
	if loopEnd >= self.position {
		self.activeLoopEnd = loopEnd
	}
	self.mutex.Unlock()
}

func (self *Looper) Length() int64 {
	self.mutex.Lock()
	var length int64
	switch streamWithLen := self.stream.(type) {
	case *bytes.Reader:
		length = int64(streamWithLen.Len())
	case StdAudioStream:
		length = streamWithLen.Length()
	default:
		panic("Looper underlying stream doesn't implement Length() int64 and is not a *bytes.Reader either")
	}
	self.mutex.Unlock()
	return length
}

func assertLoopValuesValidity(loopStart, loopEnd int64) {
	if loopStart & 0b11 != 0 { panic("loopStart must be multiple of 4") }
	if loopEnd   & 0b11 != 0 { panic("loopEnd must be multiple of 4") }
	if loopStart >= loopEnd { panic("loopStart must be strictly smaller than loopEnd") }
	if loopStart < 0 { panic("loopStart must be >= 0") }
	// Note: technically loopStart can be loopEnd - 4 or similar extremely short distances.
	//       This is allowed but it's not really correct. Nothing will sound and the looper
	//       is likely to start lagging unless absurd sample rates are used. Other small
	//       loop lengths are equally likely to cause trouble, but that's on the user.
}

// hack to be compatible with bindless
type StdAudioStream interface {
	Length() int64
	io.ReadSeeker
}

// hack to be compatible with bindless
func (self *Looper) Reset(stream io.ReadSeeker, loopStart int64, loopEnd int64) {
	self.mutex.Lock()
	self.stream = stream
	self.loopStart = loopStart
	self.loopEnd = loopEnd
	self.activeLoopEnd = loopEnd
	self.position = 0

	self.mutex.Unlock()
	pos, err := self.Seek(0, io.SeekStart)
	if err != nil { panic(err) }
	if pos != 0 { panic("non-zero position after stream reset") }
}