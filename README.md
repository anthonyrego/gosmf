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

	render.UseShader("default")
	image := render.NewSprite("box.png", 16, 16)
	for screen.IsActive() {
		image.Draw(200, 100)
		screen.BlitScreen()
	}
}
```
