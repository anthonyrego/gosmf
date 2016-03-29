package window

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/anthonyrego/dodge/input"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

// Screen object
type Screen struct {
	context     *glfw.Window
	Width       int
	Height      int
	startTime   time.Time
	frameTime   float64
	elapsedTime float64
}

// New returns a newly created Screen
func New(width int, height int, vsync bool, fullscreen bool, name string) *Screen {
	s := &Screen{}
	s.init(width, height, vsync, fullscreen, name)
	return s
}

func (window *Screen) init(width int, height int, vsync bool, fullscreen bool, name string) {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var monitor *glfw.Monitor
	if fullscreen {
		monitor = glfw.GetPrimaryMonitor()
	}
	win, err := glfw.CreateWindow(width, height, name, monitor, nil)
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

	input.AttachInputToWindow(window.context)

	window.startTime = time.Now()
	window.frameTime = time.Since(window.startTime).Seconds()
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
	time := window.GetTime().Seconds()
	window.elapsedTime = time - window.frameTime
	window.frameTime = time
}

// GetTimeSinceLastFrame will return the time since the last frame was drawn.
func (window *Screen) GetTimeSinceLastFrame() float64 {
	return window.elapsedTime
}

// GetTime returns the time duration since the start of the program
func (window *Screen) GetTime() time.Duration {
	return time.Since(window.startTime)
}

// AmountPerSecond returns the adjusted value for a per second value based on frame time
// This really needs a better function name
func (window *Screen) AmountPerSecond(persecond float64) float64 {
	return persecond * window.elapsedTime
}

// Destroy cleans up the window
func (window *Screen) Destroy() {
	window.context.Destroy()
	glfw.Terminate()
}
