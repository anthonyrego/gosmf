package input

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

// Input event types
const (
	Release int = int(glfw.Release)
	Press   int = int(glfw.Press)
	Repeat  int = int(glfw.Repeat)
)
