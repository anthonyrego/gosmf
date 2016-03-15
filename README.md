# dodge
WIP golang media framework

```
package main

import (
	"fmt"
	"github.com/anthonyrego/dodge/camera"
	"github.com/anthonyrego/dodge/input"
	"github.com/anthonyrego/dodge/render"
	"github.com/anthonyrego/dodge/timing"
	"github.com/anthonyrego/dodge/window"
	_ "image/png"
)

func main() {
	windowWidth := 800
	windowHeight := 600

	screen := window.New(windowWidth, windowHeight, true, "Dodge Example")
	defer screen.Destroy()

	render.UseShader("default")

	updateCamera := initCamera(windowWidth, windowHeight)

	image, _ := render.NewSprite("pouch.png", 264, 347)

	input.AddListener(input.KeyEscape, func(event int) {
		if event == input.Release {
			screen.SetToClose()
		}
	})

	fmt.Println("Time since start", timing.GetTime().Seconds())

	for screen.IsActive() {
		updateCamera()
		image.Draw(200, 200, 1)
		image.Draw(100, 100, 0)
		image.Draw(0, 0, 200)
		image.Draw(150, 150, 10)
		screen.BlitScreen()
	}
}

func initCamera(windowWidth int, windowHeight int) func() {
	cam1 := camera.New(true)
	cam1.SetOrtho(windowWidth, windowHeight, 200)
	cam1.SetPosition2D(0, 0)
	camx := 0.0
	camy := 0.0
	return func() {
		if input.GetKeyEventState(input.KeyA) == input.Press {
			camx -= 4
		}
		if input.GetKeyEventState(input.KeyD) == input.Press {
			camx += 4
		}
		if input.GetKeyEventState(input.KeyW) == input.Press {
			camy -= 4
		}
		if input.GetKeyEventState(input.KeyS) == input.Press {
			camy += 4
		}
		cam1.SetPosition2D(float32(camx), float32(camy))
	}
}
```
