/*
Package dodge is a work in progress simple media framework in go

A simple example that creates a window with input and draws a 2D sprite
W,A,S,D moves the 2D camera around and ESC will exit the program

   package main

   import (
   	_ "image/png"

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

   	image, _ := sprite.New("pouch.png", 264, 347)

   	font.New("VeraMono.ttf")

   	input.AddListener(input.KeyEscape, func(event int) {
   		if event == input.Release {
   			screen.SetToClose()
   		}
   	})
   	for screen.IsActive() {
   		updateCamera()
   		image.Draw(0, 0, 200)
   		image.Draw(150, 150, 10)
   		screen.BlitScreen()
   	}
   }

   func initCamera(screen *window.Screen) func() {
   	cam1 := camera.New(true)
   	cam1.SetOrtho(screen.Width*2, screen.Height*2, 200)
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
