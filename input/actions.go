package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

const (
	// Release action state for button input
	Release int = int(glfw.Release)
	// Press action state for button input
	Press int = int(glfw.Press)
	// Repeat action state for button input
	Repeat int = int(glfw.Repeat)
)
