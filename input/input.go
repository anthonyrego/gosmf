package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

var listenerList = map[int]*listener{}

type listener struct {
	callback func(event int)
}

// AddListener creates a new key listener, only the last listener for a button will be honored
func AddListener(key int, callback func(event int)) {
	listenerList[key] = &listener{callback}
}

// Callback function for input
var Callback = func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if li, found := listenerList[int(key)]; found {
		li.callback(int(action))
	}
}
