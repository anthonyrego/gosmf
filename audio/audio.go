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

// NewSound returns a newly created Sound object
// Only used if you are going to load your own data instead of using the wav loader
//
//	s := NewSound("example.mp3")
//
//  /* Code that gets the Data, BufferSize, Frequency and Channels of mp3s */
//
//  //The we load that data in manually with:
//  s.LoadPCMData()
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
	C.alBufferData(s.buffer, C.AL_FORMAT_STEREO16, unsafe.Pointer(&s.Data[0]), C.ALsizei(s.Size), C.ALsizei(s.Frequency))
}

func (s *Sound) Play() {
	var source C.ALuint
	C.alGenSources(1, &source)

	C.alSourcei(source, C.AL_BUFFER, C.ALint(s.buffer))
	C.alSourcePlay(source)
}
