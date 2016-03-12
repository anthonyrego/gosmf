package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

// KeyChannel is a channel pipe to communicate key presses to go routines
var KeyChannel = make(chan glfw.Key)

// Callback function for input
var Callback = func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		go func() {
			if len(KeyChannel) == 0 {
				KeyChannel <- key
			}
		}()
	}
}
