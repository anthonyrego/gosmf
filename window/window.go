package window

import (
	"fmt"
	"github.com/anthonyrego/dodge/input"
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
	context *glfw.Window
	Width   int
	Height  int
}

// New returns a newly created Screen
func New(width int, height int, vsync bool, name string) *Screen {
	s := &Screen{}
	s.init(width, height, vsync, name)
	return s
}

func (window *Screen) init(width int, height int, vsync bool, name string) {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()

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
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	window.context = win
	window.Width = width
	window.Height = height
	window.context.SwapBuffers()

	input.AttachInputCallbacksToWindow(window.context)
}

// IsActive returns the status of the window
func (window *Screen) IsActive() bool {
	return !window.context.ShouldClose()
}

// SetToClose will signal to close the window context
func (window *Screen) SetToClose() {
	window.context.SetShouldClose(true)
}

// BlitScreen swaps the buffers and clears the screen
func (window *Screen) BlitScreen() {
	window.context.SwapBuffers()
	glfw.PollEvents()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Destroy cleans up the window
func (window *Screen) Destroy() {
	window.context.Destroy()
	glfw.Terminate()
}
