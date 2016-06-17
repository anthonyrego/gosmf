package window

/*
#include <SDL2/SDL.h>

int getEventType(SDL_Event e) {
    return e.type;
}
int getEventKey(SDL_Event e) {
    return e.key.keysym.scancode;
}
int getEventKeyState(SDL_Event e) {
    return e.key.state;
}
*/
import "C"

const (
	WindowQuit       int = int(C.SDL_QUIT)
	KeyEventDown     int = int(C.SDL_KEYDOWN)
	KeyEventUp       int = int(C.SDL_KEYUP)
	KeyStatePressed  int = int(C.SDL_PRESSED)
	KeyStateReleased int = int(C.SDL_RELEASED)
)

func (window *Screen) runEventQueue() {
	var event C.SDL_Event

	for C.SDL_PollEvent(&event) != 0 {
		switch int(C.getEventType(event)) {
		case WindowQuit:
			window.SetToClose()
			break

		case KeyEventDown, KeyEventUp:
			if listener, found := listenerList[int(C.getEventKey(event))]; found {
				listener.callback(int(C.getEventKeyState(event)))
			}
			break
		}
	}
}
