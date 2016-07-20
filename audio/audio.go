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
	initSourceList()
}

func Cleanup() {
	destroySourceList()
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
	format := 0
	switch s.Channels {
	default:
		return
	case 1:
		format = C.AL_FORMAT_MONO16
	case 2:
		format = C.AL_FORMAT_STEREO16
	}
	C.alGenBuffers(1, &s.buffer)
	C.alBufferData(s.buffer, C.ALenum(format), unsafe.Pointer(&s.Data[0]), C.ALsizei(s.Size), C.ALsizei(s.Frequency))
}

// Play will play the sound. Returns the play request id that can be used to stop the source while playing
func (s *Sound) Play() int64 {
	source, err := requestSource()
	if err != nil {
		return -1
	}
	C.alSourcei(source.id, C.AL_SOURCE_RELATIVE, C.AL_TRUE)
	C.alSource3f(source.id, C.AL_POSITION, 0, 0, 0)
	C.alSourcei(source.id, C.AL_BUFFER, C.ALint(s.buffer))

	source.setToPlay()
	return source.requestId
}

// Play will play the sound at a given position and the falloff distance in which the sound's volume is cut in half.
// It will return the play request id that can be used to stop the source while playing
// Remember that in order for the 3D audio to work properly that the audio needs to be all in one channel, not stereo!
func (s *Sound) Play3D(x, y, z, falloff float32) int64 {
	source, err := requestSource()
	if err != nil {
		return -1
	}
	C.alSourcei(source.id, C.AL_SOURCE_RELATIVE, C.AL_FALSE)
	C.alSourcef(source.id, C.AL_REFERENCE_DISTANCE, C.ALfloat(falloff))
	C.alSource3f(source.id, C.AL_POSITION, C.ALfloat(x), C.ALfloat(y), C.ALfloat(z))
	C.alSourcei(source.id, C.AL_BUFFER, C.ALint(s.buffer))

	source.setToPlay()
	return source.requestId
}

func SetListenPosition(x, y, z float32) {
	C.alListener3f(C.AL_POSITION, C.ALfloat(x), C.ALfloat(y), C.ALfloat(z))
}
