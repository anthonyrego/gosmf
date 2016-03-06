# dodge
WIP golang media framework

```
package main

import (
	"github.com/anthonyrego/dodge/render"
	"github.com/anthonyrego/dodge/window"
	_ "image/png"
)

func main() {
	screen := window.New(800, 600, true, "Dodge Example")
	defer screen.Destroy()

	shader := render.UseShader("default")
	render.Setup2DProjection(shader, screen)

	image := render.NewSprite("box.png", 16, 16)

	for screen.IsActive() {
		image.Draw(200, 200, 1)
		image.Draw(100, 100, 0)
		image.Draw(0, 0, 200)
		image.Draw(150, 150, 10)
		screen.BlitScreen()
	}
}
```
