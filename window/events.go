package window

/*
#include <SDL2/SDL.h>

int getEventType(SDL_Event e) {
    return e.type;
}
int getEventKey(SDL_Event e) {
    return e.key.keysym.sym;
}
*/
import "C"
import (
	"fmt"
)

const (
	WindowQuit int = int(C.SDL_QUIT)
	KeyPress   int = int(C.SDL_KEYDOWN)
	KeyRelease int = int(C.SDL_KEYUP)
)

func (window *Screen) runEventLoop() {
	var event C.SDL_Event

	for C.SDL_PollEvent(&event) != 0 {
		switch int(C.getEventType(event)) {
		case WindowQuit:
			window.SetToClose()
			break

		case KeyPress, KeyRelease:
			fmt.Println("Button pressed: ", C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(C.getEventKey(event)))))
			break
		}
	}
}
