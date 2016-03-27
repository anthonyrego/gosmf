/*
Package dodge is a work in progress simple media framework in go

A simple example that creates a window with input and draws a 2D sprite
W,A,S,D moves the 2D camera around and ESC will exit the program
Enter will increase the text counter

  package main

  import (
  	_ "image/gif"
  	"math"

  	"fmt"

  	"github.com/anthonyrego/dodge/camera"
  	"github.com/anthonyrego/dodge/font"
  	"github.com/anthonyrego/dodge/input"
  	"github.com/anthonyrego/dodge/shader"
  	"github.com/anthonyrego/dodge/sprite"
  	"github.com/anthonyrego/dodge/window"
  )

  func main() {
  	windowWidth := 800
  	windowHeight := 600

  	screen := window.New(windowWidth, windowHeight, true, "Dodge Example")
  	defer screen.Destroy()

  	shader.Use("default")

  	updateCamera := initCamera(screen)
  	getCurrentFps := initFpsCounter(screen)

  	image, _ := sprite.New("sad.gif", 201, 161)

  	arial, _ := font.New("Arial.ttf")

  	buttonsPressed := arial.NewBillboard("Button Pressed 0 times", 500, 150, 10, 300)
  	elapsedTime := arial.NewBillboard("fps: ", 500, 150, 10, 300)
  	buttonCounter := 0

  	input.AddListener(input.KeyEscape, func(event int) {
  		if event == input.Release {
  			screen.SetToClose()
  		}
  	})
  	input.AddListener(input.KeyEnter, func(event int) {
  		if event == input.Press {
  			buttonCounter++
  			buttonsPressed.UpdateText(fmt.Sprintf("Button Pressed %d times", buttonCounter))
  		}
  	})

  	for screen.IsActive() {
  		updateCamera()
  		elapsedTime.UpdateText(fmt.Sprintf("fps: %d", getCurrentFps()))
  		elapsedTime.Draw(0, 500, 0)
  		buttonsPressed.Draw(0, 350, 0)
  		image.Draw(0, 0, 200)
  		screen.BlitScreen()
  	}
  }

  func initFpsCounter(screen *window.Screen) func() int {
  	const fpsBufferSize = 10
  	var fpsBuffer [fpsBufferSize]int
  	fpsCounter := 0
  	currentFps := 0

  	return func() int {
  		if fpsCounter > (fpsBufferSize - 1) {
  			fpsCounter = 0
  			fpsSum := 0
  			for i := 0; i < fpsBufferSize; i++ {
  				fpsSum += fpsBuffer[i]
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

*/
package dodge
