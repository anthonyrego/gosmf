package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/anthonyrego/gosmf/audio"
	"github.com/anthonyrego/gosmf/camera"
	"github.com/anthonyrego/gosmf/font"
	"github.com/anthonyrego/gosmf/shader"
	"github.com/anthonyrego/gosmf/sprite"
	"github.com/anthonyrego/gosmf/window"
)

func main() {

	screen := window.New(800, 600, false, 0, "gosmf example")
	defer screen.Destroy()
	defer window.Cleanup()

	audio.Init()
	defer audio.Cleanup()

	shader.Use("default")

	updateCamera := initCamera(screen)
	getCurrentFps := initFpsCounter(screen)

	// If image does not exist, download a random one
	if _, err := os.Stat("img.jpg"); os.IsNotExist(err) {
		downloadFile("https://placeimg.com/256/256/any", "img.jpg")
	}
	// If font does not exist, download roboto
	if _, err := os.Stat("Roboto-Regular.ttf"); os.IsNotExist(err) {
		downloadFile("https://raw.githubusercontent.com/google/roboto/master/hinted/Roboto-Regular.ttf",
			"Roboto-Regular.ttf")
	}

	image, _ := sprite.New("img.jpg", 256, 256)
	ttf, _ := font.New("Roboto-Regular.ttf")

	fpsDisplay := ttf.NewBillboard("fps ",
		150, 50, 6, 300, color.RGBA{255, 50, 50, 255})

	window.AddKeyListener(window.KeyEscape, func(event int) {
		if event == window.KeyStatePressed {
			screen.SetToClose()
		}
	})

	verticalSync := true
	window.AddKeyListener(window.KeyV, func(event int) {
		if event == window.KeyStateReleased {
			screen.SetVerticalSync(!verticalSync)
			verticalSync = !verticalSync
		}
	})
	cam := camera.GetActiveCamera()

	for screen.Update() {
		updateCamera()
		image.Draw(0, 0, 0, 1)
		fpsDisplay.SetText(fmt.Sprintf("fps %d", getCurrentFps()))
		fpsDisplay.Draw((cam.Position[0]+cam.Bounds[0])-float32(fpsDisplay.Width), cam.Position[1], 0)
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
		if window.GetKeyState(window.KeyA) == window.KeyStatePressed {
			camx -= screen.AmountPerSecond(speed)
		}
		if window.GetKeyState(window.KeyD) == window.KeyStatePressed {
			camx += screen.AmountPerSecond(speed)
		}
		if window.GetKeyState(window.KeyW) == window.KeyStatePressed {
			camy -= screen.AmountPerSecond(speed)
		}
		if window.GetKeyState(window.KeyS) == window.KeyStatePressed {
			camy += screen.AmountPerSecond(speed)
		}
		cam1.SetPosition2D(float32(camx), float32(camy))
	}
}

func downloadFile(url string, filename string) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(file, response.Body)
	file.Close()
}
