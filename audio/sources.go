package audio

/*
#ifdef _GOSMF_OSX_
  #include <CoreFoundation/CoreFoundation.h>
#endif

#include <OpenAL/al.h>
#include <OpenAL/alc.h>

*/
import "C"
import "fmt"

var sourceList [255]source
var sourceChannel chan int

type source struct {
	occupied  bool
	isPlaying bool
	id        C.ALuint
}

func (s *source) setToPlay() {
	C.alSourcePlay(s.id)
	s.isPlaying = true
}

func initSourceList() {
	for i, _ := range sourceList {
		C.alGenSources(1, &sourceList[i].id)
		sourceList[i].occupied = false
		sourceList[i].isPlaying = false
	}

	sourceChannel = make(chan int)
	go func() {
		var state C.ALint
		for {
			select {
			case _, ok := <-sourceChannel:
				if !ok {
					return
				}
			default:
				for i, src := range sourceList {
					if sourceList[i].occupied == true && sourceList[i].isPlaying == true {
						C.alGetSourcei(src.id, C.AL_SOURCE_STATE, &state)
						if state != C.AL_PLAYING {
							sourceList[i].occupied = false
							sourceList[i].isPlaying = false
						}
					}
				}
			}
		}
	}()
}

func destroySourceList() {
	close(sourceChannel)
	for i, _ := range sourceList {
		C.alDeleteSources(1, &sourceList[i].id)
	}
}

func requestSource() (*source, error) {
	for i, _ := range sourceList {
		if !sourceList[i].occupied {
			sourceList[i].occupied = true
			return &sourceList[i], nil
		}
	}
	return nil, fmt.Errorf("no available sources")
}
