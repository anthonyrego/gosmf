package audio

/*
#cgo LDFLAGS: -framework OpenAL
#include <CoreFoundation/CoreFoundation.h>
#include <AudioToolbox/AudioToolbox.h>
#import <OpenAL/al.h>
#import <OpenAL/alc.h>

ALCcontext *context;
ALCdevice *device;

void openAudioDevice() {
  const ALCchar *default_device;

  default_device = alcGetString(NULL, ALC_DEFAULT_DEVICE_SPECIFIER);
  if ((device = alcOpenDevice(default_device)) == NULL)
  {
    fprintf(stderr, "failed to open sound device\n");
    return;
  }
  context = alcCreateContext(device, NULL);

  alcMakeContextCurrent(context);

  alcProcessContext(context);
  alGetError();
}

void closeAudioDevice()
{
  alcMakeContextCurrent(NULL);
  alcDestroyContext(context);
  alcCloseDevice(device);
}
*/
import "C"

func Init() {
	C.openAudioDevice()
}

func Cleanup() {
	C.closeAudioDevice()
}
