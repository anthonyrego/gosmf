# dodge
WIP golang media framework

```
package main

import (
	"github.com/anthonyrego/dodge/texture"
	"github.com/anthonyrego/dodge/window"
	_ "image/png"
)

func main() {
	screen := window.New(1024, 768, true, "Dodge Example")
	defer screen.Destroy()

	texture.New("box.png")
	for screen.IsActive() {
		screen.BlitScreen()
	}
}
```
