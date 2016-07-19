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

type Sound struct {
	channels  uint16
	frequency uint32
	size      uint32
	data      []byte
	buffer    C.ALuint
}

func (s *Sound) attachSoundData() {
	C.alGenBuffers(1, &s.buffer)
	C.alBufferData(s.buffer, C.AL_FORMAT_STEREO16, unsafe.Pointer(&s.data[0]), C.ALsizei(s.size), C.ALsizei(s.frequency))
}

func (s *Sound) Play() {
	var source C.ALuint
	C.alGenSources(1, &source)

	C.alSourcei(source, C.AL_BUFFER, C.ALint(s.buffer))
	C.alSourcePlay(source)
}
