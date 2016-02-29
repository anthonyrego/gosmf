# dodge
WIP golang media framework

```
package main

import (
	"github.com/anthonyrego/dodge/render"
	"github.com/anthonyrego/dodge/texture"
	"github.com/anthonyrego/dodge/window"
	_ "image/png"
)

func main() {
	screen := window.New(600, 600, true, "Dodge Example")
	defer screen.Destroy()

	tex := texture.New("box.png")
	render.UseShader(&render.DefaultShader)
	for screen.IsActive() {
		render.DrawSprite(tex, 128, 128, 0, 0)
		screen.BlitScreen()
	}
}
```
