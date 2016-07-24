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
	Channels      uint16
	Frequency     uint32
	BitsPerSample uint16
	Size          uint32
	Data          []byte
	buffer        C.ALuint
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

	if s.Channels > 1 {
		switch s.BitsPerSample {
		case 8:
			format = C.AL_FORMAT_STEREO8
		case 16:
			format = C.AL_FORMAT_STEREO16
		}
	} else {
		switch s.BitsPerSample {
		case 8:
			format = C.AL_FORMAT_MONO8
		case 16:
			format = C.AL_FORMAT_MONO16
		}
	}

	C.alGenBuffers(1, &s.buffer)
	C.alBufferData(s.buffer, C.ALenum(format), unsafe.Pointer(&s.Data[0]), C.ALsizei(s.Size), C.ALsizei(s.Frequency))
}

func (s *Sound) Destroy() {
	C.alDeleteBuffers(1, &s.buffer)
}

// Play will play the sound. Volume ( 1.0 is normal volume, 0 is silence )
// Returns the PlayInstance that can be used to stop the source while playing
func (s *Sound) Play(volume float32) (request PlayInstance) {
	source, err := requestSource()
	if err != nil {
		return request
	}
	C.alSourcef(source.id, C.AL_GAIN, C.ALfloat(volume))
	C.alSourcei(source.id, C.AL_SOURCE_RELATIVE, C.AL_TRUE)
	C.alSource3f(source.id, C.AL_POSITION, 0, 0, 0)
	C.alSourcei(source.id, C.AL_BUFFER, C.ALint(s.buffer))

	source.setToPlay()
	request.id = source.requestId
	request.src = source
	return request
}

// Play will play the sound at a given position, the falloff distance in which the sound's volume is cut in half,
// and the volume ( 1.0 is normal volume, 0 is silence )
// It will return the PlayInstance that can be used to stop the source while playing
// Remember that in order for the 3D audio to work properly that the audio needs to be all in one channel, not stereo!
func (s *Sound) Play3D(x, y, z, falloff, volume float32) (request PlayInstance) {
	source, err := requestSource()
	if err != nil {
		return request
	}
	C.alSourcef(source.id, C.AL_GAIN, C.ALfloat(volume))
	C.alSourcei(source.id, C.AL_SOURCE_RELATIVE, C.AL_FALSE)
	C.alSourcef(source.id, C.AL_REFERENCE_DISTANCE, C.ALfloat(falloff))
	C.alSource3f(source.id, C.AL_POSITION, C.ALfloat(x), C.ALfloat(y), C.ALfloat(z))
	C.alSourcei(source.id, C.AL_BUFFER, C.ALint(s.buffer))

	source.setToPlay()
	request.id = source.requestId
	request.src = source
	return request
}

// PlayInstance is returned when you make a call to play a sound so you can stop playback or determine if the sound is still playing
type PlayInstance struct {
	src *source
	id  int64
}

func (playback *PlayInstance) StopPlayback() {
	if playback.src != nil &&
		playback.id == playback.src.requestId &&
		playback.src.isPlaying {

		C.alSourceStop(playback.src.id)
		playback.src.occupied = false
		playback.src.isPlaying = false
	}
}

func (playback *PlayInstance) IsPlaying() bool {
	if playback.src != nil &&
		playback.id == playback.src.requestId &&
		playback.src.isPlaying {
		return true
	}
	return false
}

func SetListenPosition(x, y, z float32) {
	C.alListener3f(C.AL_POSITION, C.ALfloat(x), C.ALfloat(y), C.ALfloat(z))
}
