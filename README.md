# dodge
WIP golang media framework

```
package main

import (
	"github.com/anthonyrego/dodge/window"
)

func main() {
	screen := new(window.Screen)
	screen.Init(1024, 768, "Dodge Example")
	defer screen.Destroy()

	for screen.IsActive() {
		screen.BlitScreen()
	}
}
```
