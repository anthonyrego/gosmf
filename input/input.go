package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

var listenerList = map[int]*listener{}

type listener struct {
	callback func(action int)
}

// AddListener creates a new key listener
func AddListener(key int, callback func(action int)) {
	listenerList[key] = &listener{callback}
}

// Callback function for input
var Callback = func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if li, found := listenerList[int(key)]; found {
		li.callback(int(action))
	}
}
