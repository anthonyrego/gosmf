package window

/*
#include <SDL2/SDL.h>
*/
import "C"

// Remapped sdl keys so we dont need import C all the time
const (
	KeyUnknown    int = int(C.SDL_SCANCODE_UNKNOWN)
	KeyBackspace  int = int(C.SDL_SCANCODE_BACKSPACE)
	KeyTab        int = int(C.SDL_SCANCODE_TAB)
	KeyClear      int = int(C.SDL_SCANCODE_CLEAR)
	KeyReturn     int = int(C.SDL_SCANCODE_RETURN)
	KeyPause      int = int(C.SDL_SCANCODE_PAUSE)
	KeyEscape     int = int(C.SDL_SCANCODE_ESCAPE)
	KeySpace      int = int(C.SDL_SCANCODE_SPACE)
	KeyComma      int = int(C.SDL_SCANCODE_COMMA)
	KeyMinus      int = int(C.SDL_SCANCODE_MINUS)
	KeyPeriod     int = int(C.SDL_SCANCODE_PERIOD)
	KeySlash      int = int(C.SDL_SCANCODE_SLASH)
	Key0          int = int(C.SDL_SCANCODE_0)
	Key1          int = int(C.SDL_SCANCODE_1)
	Key2          int = int(C.SDL_SCANCODE_2)
	Key3          int = int(C.SDL_SCANCODE_3)
	Key4          int = int(C.SDL_SCANCODE_4)
	Key5          int = int(C.SDL_SCANCODE_5)
	Key6          int = int(C.SDL_SCANCODE_6)
	Key7          int = int(C.SDL_SCANCODE_7)
	Key8          int = int(C.SDL_SCANCODE_8)
	Key9          int = int(C.SDL_SCANCODE_9)
	KeySemicolon  int = int(C.SDL_SCANCODE_SEMICOLON)
	KeyEquals     int = int(C.SDL_SCANCODE_EQUALS)
	KeyLBracket   int = int(C.SDL_SCANCODE_LEFTBRACKET)
	KeyBackslash  int = int(C.SDL_SCANCODE_BACKSLASH)
	KeyRBracket   int = int(C.SDL_SCANCODE_RIGHTBRACKET)
	KeyA          int = int(C.SDL_SCANCODE_A)
	KeyB          int = int(C.SDL_SCANCODE_B)
	KeyC          int = int(C.SDL_SCANCODE_C)
	KeyD          int = int(C.SDL_SCANCODE_D)
	KeyE          int = int(C.SDL_SCANCODE_E)
	KeyF          int = int(C.SDL_SCANCODE_F)
	KeyG          int = int(C.SDL_SCANCODE_G)
	KeyH          int = int(C.SDL_SCANCODE_H)
	KeyI          int = int(C.SDL_SCANCODE_I)
	KeyJ          int = int(C.SDL_SCANCODE_J)
	KeyK          int = int(C.SDL_SCANCODE_K)
	KeyL          int = int(C.SDL_SCANCODE_L)
	KeyM          int = int(C.SDL_SCANCODE_M)
	KeyN          int = int(C.SDL_SCANCODE_N)
	KeyO          int = int(C.SDL_SCANCODE_O)
	KeyP          int = int(C.SDL_SCANCODE_P)
	KeyQ          int = int(C.SDL_SCANCODE_Q)
	KeyR          int = int(C.SDL_SCANCODE_R)
	KeyS          int = int(C.SDL_SCANCODE_S)
	KeyT          int = int(C.SDL_SCANCODE_T)
	KeyU          int = int(C.SDL_SCANCODE_U)
	KeyV          int = int(C.SDL_SCANCODE_V)
	KeyW          int = int(C.SDL_SCANCODE_W)
	KeyX          int = int(C.SDL_SCANCODE_X)
	KeyY          int = int(C.SDL_SCANCODE_Y)
	KeyZ          int = int(C.SDL_SCANCODE_Z)
	KeyDelete     int = int(C.SDL_SCANCODE_DELETE)
	KeyKP0        int = int(C.SDL_SCANCODE_KP_0)
	KeyKP1        int = int(C.SDL_SCANCODE_KP_1)
	KeyKP2        int = int(C.SDL_SCANCODE_KP_2)
	KeyKP3        int = int(C.SDL_SCANCODE_KP_3)
	KeyKP4        int = int(C.SDL_SCANCODE_KP_4)
	KeyKP5        int = int(C.SDL_SCANCODE_KP_5)
	KeyKP6        int = int(C.SDL_SCANCODE_KP_6)
	KeyKP7        int = int(C.SDL_SCANCODE_KP_7)
	KeyKP8        int = int(C.SDL_SCANCODE_KP_8)
	KeyKP9        int = int(C.SDL_SCANCODE_KP_9)
	KeyKPPeriod   int = int(C.SDL_SCANCODE_KP_PERIOD)
	KeyKPDivide   int = int(C.SDL_SCANCODE_KP_DIVIDE)
	KeyKPMultiply int = int(C.SDL_SCANCODE_KP_MULTIPLY)
	KeyKPMinus    int = int(C.SDL_SCANCODE_KP_MINUS)
	KeyKPPlus     int = int(C.SDL_SCANCODE_KP_PLUS)
	KeyKPEnter    int = int(C.SDL_SCANCODE_KP_ENTER)
	KeyKPEquals   int = int(C.SDL_SCANCODE_KP_EQUALS)
	KeyUp         int = int(C.SDL_SCANCODE_UP)
	KeyDown       int = int(C.SDL_SCANCODE_DOWN)
	KeyRight      int = int(C.SDL_SCANCODE_RIGHT)
	KeyLeft       int = int(C.SDL_SCANCODE_LEFT)
	KeyInsert     int = int(C.SDL_SCANCODE_INSERT)
	KeyHome       int = int(C.SDL_SCANCODE_HOME)
	KeyEnd        int = int(C.SDL_SCANCODE_END)
	KeyPageUp     int = int(C.SDL_SCANCODE_PAGEUP)
	KeyPageDown   int = int(C.SDL_SCANCODE_PAGEDOWN)
	KeyF1         int = int(C.SDL_SCANCODE_F1)
	KeyF2         int = int(C.SDL_SCANCODE_F2)
	KeyF3         int = int(C.SDL_SCANCODE_F3)
	KeyF4         int = int(C.SDL_SCANCODE_F4)
	KeyF5         int = int(C.SDL_SCANCODE_F5)
	KeyF6         int = int(C.SDL_SCANCODE_F6)
	KeyF7         int = int(C.SDL_SCANCODE_F7)
	KeyF8         int = int(C.SDL_SCANCODE_F8)
	KeyF9         int = int(C.SDL_SCANCODE_F9)
	KeyF10        int = int(C.SDL_SCANCODE_F10)
	KeyF11        int = int(C.SDL_SCANCODE_F11)
	KeyF12        int = int(C.SDL_SCANCODE_F12)
	KeyF13        int = int(C.SDL_SCANCODE_F13)
	KeyF14        int = int(C.SDL_SCANCODE_F14)
	KeyF15        int = int(C.SDL_SCANCODE_F15)
	KeyCapsLock   int = int(C.SDL_SCANCODE_CAPSLOCK)
	KeyRShift     int = int(C.SDL_SCANCODE_RSHIFT)
	KeyLShift     int = int(C.SDL_SCANCODE_LSHIFT)
	KeyRControl   int = int(C.SDL_SCANCODE_RCTRL)
	KeyLControl   int = int(C.SDL_SCANCODE_LCTRL)
	KeyRAlt       int = int(C.SDL_SCANCODE_RALT)
	KeyLAlt       int = int(C.SDL_SCANCODE_LALT)
	KeyMode       int = int(C.SDL_SCANCODE_MODE)
	KeyHelp       int = int(C.SDL_SCANCODE_HELP)
	KeySysreq     int = int(C.SDL_SCANCODE_SYSREQ)
	KeyMenu       int = int(C.SDL_SCANCODE_MENU)
	KeyPower      int = int(C.SDL_SCANCODE_POWER)
	KeyUndo       int = int(C.SDL_SCANCODE_UNDO)
)
