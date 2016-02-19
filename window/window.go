package window

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

// Screen object
type Screen struct {
	Context *glfw.Window
}

// Init the window object
func (window *Screen) Init(width int, height int, vsync bool, name string) {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window.Context, err = glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		panic(err)
	}
	window.Context.MakeContextCurrent()

	if vsync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	window.Context.SwapBuffers()
}

// IsActive returns the status of the window
func (window *Screen) IsActive() bool {
	return !window.Context.ShouldClose()
}

// BlitScreen swaps the buffers and clears the screen
func (window *Screen) BlitScreen() {
	window.Context.SwapBuffers()
	glfw.PollEvents()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Destroy cleans up the window
func (window *Screen) Destroy() {
	window.Context.Destroy()
	glfw.Terminate()
}
