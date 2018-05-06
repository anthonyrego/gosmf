package window

/*
#include <SDL2/SDL.h>

typedef struct MOUSESTATE {
	int X;
	int Y;
	int State;
}mouse;

int getKeyState(int key) {
  const Uint8 *state = SDL_GetKeyboardState(NULL);

  if (state[key]){
    return SDL_PRESSED;
  }
  return SDL_RELEASED;
}

char *getClipboard() {
  if (SDL_HasClipboardText() == SDL_TRUE){
    return SDL_GetClipboardText();
  }
  return "";
}

mouse getMouseState() {
	mouse m;
	m.State = SDL_GetMouseState(&m.X, &m.Y);
	return m;
}

mouse getRelativeMouseState() {
	mouse m;
	m.State = SDL_GetRelativeMouseState(&m.X, &m.Y);
	return m;
}

*/
import "C"

type mouseState C.struct_MOUSESTATE

var listenerList = map[int]*listener{}

type listener struct {
	callback func(event int)
}

// AddKeyListener creates a new key listener, only the last listener for a button will be honored
//	input.AddKeyListener(input.KeyEscape, func(event int) {
//		if event == input.KeyStateReleased {
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

// GetMouseState returns the state and position of the mouse
func GetMouseState() (int, int, int) {
	m := mouseState(C.getMouseState())
	return int(m.X), int(m.Y), int(m.State)
}

// GetRelativeMouseState returns the state the position of the mouse
// relative to when this function was last called
func GetRelativeMouseState() (int, int, int) {
	m := mouseState(C.getRelativeMouseState())
	return int(m.X), int(m.Y), int(m.State)
}

// GetClipboard will return text from the clipboard
func GetClipboard() string {
	return C.GoString(C.getClipboard())
}
