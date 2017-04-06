package main

import (
	"fmt"

	//"github.com/SIRHAMY/Hypno/hypno"
	"github.com/SIRHAMY/Hypno/inputs/audio"
)

func main() {
	fmt.Println("hello world")

	var ss audio.SoundService
	var s = ss.NewSound()

	fmt.Print(s.Name)
}
