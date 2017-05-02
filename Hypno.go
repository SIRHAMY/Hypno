package main

import (
	"fmt"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {
	// fmt.Println("hello world")

	// var soundService audio.SoundService
	// var sound = soundService.NewSound()

	// fmt.Println(sound.Name)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int32, 8192)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()

	for true {
		time.Sleep(2 * time.Second)
		fmt.Println("Input latency: " + stream.Info().InputLatency.String())
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
