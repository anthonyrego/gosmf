package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/anthonyrego/gosmf/audio"
)

func main() {

	audio.Init()
	defer audio.Cleanup()

	// If sound does not exist, download one
	if _, err := os.Stat("img.jpg"); os.IsNotExist(err) {
		downloadFile("https://archive.org/download/Sound_Effects_3/DOORBELL.WAV", "door.wav")
	}
	sound := audio.LoadWav("door.wav")
	playRequest := sound.Play3D(0, 0, 0, 100)

	for audio.IsPlaying(playRequest) {
		fmt.Print("\rPlaying...")
	}
	fmt.Println("done")
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
