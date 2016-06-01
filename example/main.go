package main

import (
	"image/color"
	_ "image/png"
	"math"

	"fmt"

	"github.com/anthonyrego/gosmf/camera"
	"github.com/anthonyrego/gosmf/font"
	"github.com/anthonyrego/gosmf/input"
	"github.com/anthonyrego/gosmf/shader"
	"github.com/anthonyrego/gosmf/sprite"
	"github.com/anthonyrego/gosmf/window"
)

func main() {
	windowWidth := 800
	windowHeight := 600

	screen := window.New(windowWidth, windowHeight, false, "gosmf example")
	defer window.Cleanup()

	shader.Use("default")

	updateCamera := initCamera(screen)
	getCurrentFps := initFpsCounter(screen)

	image, _ := sprite.New("box.png", 16, 16)

	ttf, _ := font.New("Roboto-Regular.ttf")

	buttonsPressed := ttf.NewBillboard("Button Pressed 0 times",
		500, 150, 2, 20, 300, color.RGBA{255, 255, 255, 255})
	fpsDisplay := ttf.NewBillboard("fps: ",
		500, 250, 2, 64, 300, color.RGBA{255, 255, 255, 255})

	input.AddListener(input.KeyEscape, func(event int) {
		if event == input.Release {
			screen.SetToClose()
		}
	})

	verticalSync := true
	input.AddListener(input.KeyV, func(event int) {
		if event == input.Release {
			screen.SetVerticalSync(!verticalSync)
			verticalSync = !verticalSync
		}
	})

	buttonCounter := 0
	input.AddListener(input.KeyEnter, func(event int) {
		if event == input.Press {
			buttonCounter++
			buttonsPressed.SetText(fmt.Sprintf("Button Pressed %d times", buttonCounter))
		}
	})

	for screen.IsActive() {
		updateCamera()
		fpsDisplay.SetText(fmt.Sprintf("fps: %d", getCurrentFps()))
		fpsDisplay.Draw(0, 300, 0)
		buttonsPressed.Draw(0, 200, 0)
		image.Draw(0, 500, 0, 50)
		screen.BlitScreen()
	}
}

func initFpsCounter(screen *window.Screen) func() int {
	const fpsBufferSize = 20
	var fpsBuffer [fpsBufferSize]int
	fpsCounter := 0
	currentFps := 0

	return func() int {
		if fpsCounter > (fpsBufferSize - 1) {
			fpsCounter = 0
			fpsSum := 0
			for _, val := range fpsBuffer {
				fpsSum += val
			}
			currentFps = int(math.Ceil((float64(fpsSum) / float64(fpsBufferSize))))
			if currentFps < 0 {
				currentFps = 0
			}
		}

		fpsBuffer[fpsCounter] = int(1 / screen.GetTimeSinceLastFrame())
		fpsCounter++
		return currentFps
	}
}

func initCamera(screen *window.Screen) func() {
	cam1 := camera.New(true)
	cam1.SetOrtho(screen.Width, screen.Height, 200)
	cam1.SetPosition2D(0, 0)
	camx := 0.0
	camy := 0.0
	speed := 300.0
	return func() {
		if input.GetKeyEventState(input.KeyA) == input.Press {
			camx -= screen.AmountPerSecond(speed)
		}
		if input.GetKeyEventState(input.KeyD) == input.Press {
			camx += screen.AmountPerSecond(speed)
		}
		if input.GetKeyEventState(input.KeyW) == input.Press {
			camy -= screen.AmountPerSecond(speed)
		}
		if input.GetKeyEventState(input.KeyS) == input.Press {
			camy += screen.AmountPerSecond(speed)
		}
		cam1.SetPosition2D(float32(camx), float32(camy))
	}
}
