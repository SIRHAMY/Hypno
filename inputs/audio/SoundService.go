package audio

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

type SoundService struct {
	Name        string
	soundStream *portaudio.Stream
}

type Sound struct {
	Name string
}

func (s SoundService) InputInfo() string {
	return s.Name
}

func (s SoundService) NewSound() *Sound {
	var sound Sound
	sound.Name = "HAMY Song"
	fmt.Println("YoYOYO")
	return &sound
}
