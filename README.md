# dodge
WIP golang media framework

```
package main

import (
	"fmt"
	"github.com/anthonyrego/dodge/camera"
	"github.com/anthonyrego/dodge/input"
	"github.com/anthonyrego/dodge/render"
	"github.com/anthonyrego/dodge/window"
	_ "image/png"
)

func main() {
	windowWidth := 800
	windowHeight := 600

	screen := window.New(windowWidth, windowHeight, true, "Dodge Example")
	defer screen.Destroy()

	render.UseShader("default")

	cam1 := camera.New(true)
	cam1.SetOrtho(windowWidth, windowHeight, 200)
	cam1.SetPosition2D(0, 0)

	image, _ := render.NewSprite("pouch.png", 264, 347)

	go func() {
		for {
			key := <-input.KeyChannel
			fmt.Println("Button pressed ", key)
		}
	}()

	for screen.IsActive() {
		image.Draw(200, 200, 1)
		image.Draw(100, 100, 0)
		image.Draw(0, 0, 200)
		image.Draw(150, 150, 10)
		screen.BlitScreen()
	}
}
```
