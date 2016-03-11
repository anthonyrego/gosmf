package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

// Callback function for input
var Callback = func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Release {
		window.SetShouldClose(true)
	}
}
