package window

/*
#include <SDL2/SDL.h>

int getKeyState(SDL_Keycode key) {
  const Uint8 *state = SDL_GetKeyboardState(NULL);
  SDL_Scancode scancode = SDL_GetScancodeFromKey(key);

  if (state[scancode]){
    return SDL_PRESSED;
  }
  return SDL_RELEASED;
}
*/
import "C"

var listenerList = map[int]*listener{}

type listener struct {
	callback func(event int)
}

// AddListener creates a new key listener, only the last listener for a button will be honored
//	input.AddListener(input.KeyEscape, func(event int) {
//		if event == input.Release {
//			fmt.Println("Escape button released!")
//		}
//	})
func AddListener(key int, callback func(event int)) {
	listenerList[key] = &listener{callback}
}

// DestroyListener removes listener for a key
func DestroyListener(key int) {
	if _, ok := listenerList[key]; ok {
		listenerList[key].callback = func(event int) {}
	}
}

// GetKeyState will return the event state for a key
func GetKeyState(key int) int {
	return int(C.getKeyState(C.SDL_Keycode(key)))
}
