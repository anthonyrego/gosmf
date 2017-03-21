package window

/*
#include <SDL2/SDL.h>

int getEventType(SDL_Event e) {
    return e.type;
}

int getWindowEventType(SDL_Event e) {
    return e.window.event;
}

int getEventKey(SDL_Event e) {
    return e.key.keysym.scancode;
}
int getEventKeyState(SDL_Event e) {
    return e.key.state;
}
SDL_TextInputEvent getInputText(SDL_Event e) {
  return e.text;
}
*/
import "C"

const (
	WindowQuit       int = int(C.SDL_QUIT)
	KeyEventDown     int = int(C.SDL_KEYDOWN)
	KeyEventUp       int = int(C.SDL_KEYUP)
	KeyStatePressed  int = int(C.SDL_PRESSED)
	KeyStateReleased int = int(C.SDL_RELEASED)
	TextInput        int = int(C.SDL_TEXTINPUT)
	WindowEvent      int = int(C.SDL_WINDOWEVENT)
	WindowResized    int = int(C.SDL_WINDOWEVENT_SIZE_CHANGED)
)

var inputTextCallback = func(text string) {}

func SetInputCallback(callback func(text string)) {
	inputTextCallback = callback
}

func UnSetInputCallback() {
	inputTextCallback = func(text string) {}
}

func (window *Screen) SetResizeCallback(callback func(w, h int)) {
	window.resizedCallback = callback
}

func (window *Screen) UnSetResizeCallback() {
	window.resizedCallback = func(w, h int) {}
}

func (window *Screen) runEventQueue() {
	var event C.SDL_Event

	for C.SDL_PollEvent(&event) != 0 {
		switch int(C.getEventType(event)) {
		case WindowQuit:
			window.SetToClose()
			break

		case WindowEvent:
			switch int(C.getWindowEventType(event)) {
			case WindowResized:
				w := C.int(window.Width)
				h := C.int(window.Height)
				C.SDL_GL_GetDrawableSize(window.sdlWindow, &w, &h)
				window.Width = int(w)
				window.Height = int(h)
				window.resizedCallback(window.Width, window.Height)
				break
			}

		case KeyEventDown, KeyEventUp:
			if listener, found := listenerList[int(C.getEventKey(event))]; found {
				listener.callback(int(C.getEventKeyState(event)))
			}
			break

		case TextInput:
			ev := C.getInputText(event)
			inputTextCallback(C.GoString(&ev.text[0]))
			break
		}
	}
}
