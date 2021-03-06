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
	if _, err := os.Stat("door.wav"); os.IsNotExist(err) {
		downloadFile("https://archive.org/download/Sound_Effects_3/DOORBELL.WAV", "door.wav")
	}
	sound := audio.LoadWav("door.wav")
	playInstance := sound.Play(1)

	for playInstance.IsPlaying() {
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
