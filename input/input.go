package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

var listenerList = map[int]*listener{}
var window *glfw.Window

type listener struct {
	callback func(event int)
}

// AddListener creates a new key listener, only the last listener for a button will be honored
//	input.AddListener(input.KeyEscape, func(event int) {
//		if event == input.Release {
//			fmt.Println("Escape button released!")
//		}
//	})
func AddListener(key int, callback func(event int)) {
	listenerList[key] = &listener{callback}
}

// AttachInputToWindow will enable the input callbacks on specified window
func AttachInputToWindow(win *glfw.Window) {
	win.SetKeyCallback(callback)
	window = win
}

// GetKeyEventState will return the event state for a key
func GetKeyEventState(key int) int {
	if window != nil {
		return int(window.GetKey(glfw.Key(key)))
	}
	return 0
}

var callback = func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if li, found := listenerList[int(key)]; found {
		li.callback(int(action))
	}
}
