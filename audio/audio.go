package audio

/*
#ifdef _GOSMF_OSX_
  #include <CoreFoundation/CoreFoundation.h>
#endif

#include <OpenAL/al.h>
#include <OpenAL/alc.h>


ALuint genBuffer() {
  ALuint sound;

  alGenBuffers(1, &sound);
  return sound;
}

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

func (w *Sound) attachSoundData() {
	w.buffer = C.genBuffer()
	C.alBufferData(w.buffer, C.AL_FORMAT_STEREO16, unsafe.Pointer(&w.data[0]), C.ALsizei(w.size), C.ALsizei(w.frequency))
}

func (w *Sound) Play() {
	var source C.ALuint
	C.alGenSources(1, &source)

	C.alSourcei(source, C.AL_BUFFER, C.ALint(w.buffer))
	C.alSourcePlay(source)
}
