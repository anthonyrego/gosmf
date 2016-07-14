package window

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
)

/*
#include <SDL2/SDL.h>

void setGlContextAttributes() {
	SDL_GL_SetAttribute(SDL_GL_CONTEXT_PROFILE_MASK, SDL_GL_CONTEXT_PROFILE_CORE);
	SDL_GL_SetAttribute(SDL_GL_CONTEXT_FLAGS, SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG);
	#ifndef _GOSMF_OSX_
		SDL_GL_SetAttribute(SDL_GL_CONTEXT_MAJOR_VERSION, 4);
		SDL_GL_SetAttribute(SDL_GL_CONTEXT_MINOR_VERSION, 1);
	#endif
}
*/
import "C"

func init() {
	runtime.LockOSThread()
}

// Screen object
type Screen struct {
	sdlWindow   *C.SDL_Window
	renderer    *C.SDL_Renderer
	Width       int
	Height      int
	startTime   time.Time
	elapsedTime float64
	frameTime   time.Time
	name        string
	shouldClose bool
	vsync       bool
}

// New returns a newly created Screen
func New(width int, height int, fullscreen bool, FSAA int, name string) *Screen {
	window := &Screen{}

	C.SDL_Init(C.SDL_INIT_VIDEO)
	C.setGlContextAttributes()

	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)

	// Force hardware accel
	C.SDL_GL_SetAttribute(C.SDL_GL_ACCELERATED_VISUAL, 1)

	if FSAA > 0 {
		// FSAA (Fullscreen antialiasing)
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, C.int(FSAA)) // 2, 4, 8
	}

	flags := C.SDL_WINDOW_OPENGL | C.SDL_RENDERER_ACCELERATED
	if fullscreen {
		flags = flags | C.SDL_WINDOW_FULLSCREEN
	}

	C.SDL_CreateWindowAndRenderer(C.int(width), C.int(height), C.Uint32(flags), &window.sdlWindow, &window.renderer)
	C.SDL_SetWindowTitle(window.sdlWindow, C.CString(name))
	C.SDL_GL_CreateContext(window.sdlWindow)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure global settings
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	window.Width = width
	window.Height = height
	window.name = name
	window.shouldClose = false
	C.SDL_GL_SwapWindow(window.sdlWindow)

	window.startTime = time.Now()
	window.frameTime = time.Now()
	C.SDL_GL_SetSwapInterval(1)
	window.vsync = true

	return window
}

// IsActive returns the status of the window
func (window *Screen) IsActive() bool {
	return !window.shouldClose
}

// Update will draw the new frame, run the event queue and check if the window
// is still active. It will return the window active state.
func (window *Screen) Update() bool {
	window.runEventQueue()
	window.blitScreen()
	return window.IsActive()
}

// SetToClose will signal to close the window context
func (window *Screen) SetToClose() {
	window.shouldClose = true
}

// blitScreen swaps the buffers and clears the screen
func (window *Screen) blitScreen() {
	C.SDL_GL_SwapWindow(window.sdlWindow)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	window.elapsedTime = time.Since(window.frameTime).Seconds()
	window.frameTime = time.Now()
}

// GetTimeSinceLastFrame will return the time since the last frame was drawn.
func (window *Screen) GetTimeSinceLastFrame() float64 {
	return window.elapsedTime
}

// GetTime returns the time duration since the start of the program
func (window *Screen) GetTime() time.Duration {
	return time.Since(window.startTime)
}

// SetClearColor sets the color when the screen is cleared for redrawing
func (window *Screen) SetClearColor(r float32, g float32, b float32, a float32) {
	gl.ClearColor(r, g, b, a)
}

// SetVerticalSync sets the vertical sync status
func (window *Screen) SetVerticalSync(enabled bool) {
	if enabled {
		C.SDL_GL_SetSwapInterval(1)
	} else {
		C.SDL_GL_SetSwapInterval(0)
	}
	window.vsync = enabled
}

// AmountPerSecond returns the adjusted value for a per second value based on frame time
// This really needs a better function name
func (window *Screen) AmountPerSecond(persecond float64) float64 {
	return persecond * window.elapsedTime
}

// Destroy cleans up the window
func (window *Screen) Destroy() {
	C.SDL_DestroyRenderer(window.renderer)
	C.SDL_DestroyWindow(window.sdlWindow)
}

// Cleanup should be called before the program closes
func Cleanup() {
	C.SDL_Quit()
}
