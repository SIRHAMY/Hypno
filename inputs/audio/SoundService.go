package audio

import "github.com/SIRHAMY/Hypno/hypno"

type SoundService struct {
}

func (s *SoundService) NewSound() *hypno.Sound {
	var sound hypno.Sound
	sound.Name = "HAMY Song"
	return &sound
}
