package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

const (
	// Release event state for button input
	Release int = int(glfw.Release)
	// Press event state for button input
	Press int = int(glfw.Press)
	// Repeat event state for button input
	Repeat int = int(glfw.Repeat)
)
