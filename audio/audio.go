package audio

/*
#ifdef _GOSMF_OSX_
  #include <CoreFoundation/CoreFoundation.h>
#endif

#include <OpenAL/al.h>
#include <OpenAL/alc.h>

*/
import "C"
import "unsafe"

func Init() {
	initDevice()
}

func Cleanup() {
	destroyDevice()
}

var soundList = map[string]*Sound{}

type Sound struct {
	Channels  uint16
	Frequency uint32
	Size      uint32
	Data      []byte
	buffer    C.ALuint
}

// NewSound returns a newly created Sound object.
//
// Only used if you are going to load your own data instead of using the wav loader
//
//	s := NewSound("example.mp3")
//
//	/* Code that gets the Data, BufferSize, Frequency and Channels of mp3s */
//
//	//The we load that data in manually with:
//	s.LoadPCMData()
//
func NewSound(file string) *Sound {

	if sound, found := soundList[file]; found {
		return sound
	}

	s := &Sound{}
	soundList[file] = s
	return soundList[file]
}

func (s *Sound) LoadPCMData() {
	C.alGenBuffers(1, &s.buffer)
	C.alBufferData(s.buffer, C.AL_FORMAT_MONO16, unsafe.Pointer(&s.Data[0]), C.ALsizei(s.Size), C.ALsizei(s.Frequency))
}

func (s *Sound) Play(x, y, z float32) {
	var source C.ALuint
	C.alGenSources(1, &source)
	C.alSourcef(source, C.AL_REFERENCE_DISTANCE, 100)
	C.alSource3f(source, C.AL_POSITION, C.ALfloat(x), C.ALfloat(y), C.ALfloat(z))
	C.alSourcei(source, C.AL_BUFFER, C.ALint(s.buffer))
	C.alSourcePlay(source)
}

func SetListenPosition(x, y, z float32) {
	C.alListenerf(C.AL_GAIN, 1)
	C.alListener3f(C.AL_POSITION, C.ALfloat(x), C.ALfloat(y), C.ALfloat(z))
}
