package window

/*
#include <SDL2/SDL.h>

int getKeyState(int key) {
  const Uint8 *state = SDL_GetKeyboardState(NULL);

  if (state[key]){
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

// AddKeyListener creates a new key listener, only the last listener for a button will be honored
//	input.AddListener(input.KeyEscape, func(event int) {
//		if event == input.Release {
//			fmt.Println("Escape button released!")
//		}
//	})
func AddKeyListener(key int, callback func(event int)) {
	listenerList[key] = &listener{callback}
}

// DestroyKeyListener removes listener for a key
func DestroyKeyListener(key int) {
	if _, ok := listenerList[key]; ok {
		listenerList[key].callback = func(event int) {}
	}
}

// GetKeyState will return the event state for a key
func GetKeyState(key int) int {
	return int(C.getKeyState(C.int(key)))
}
