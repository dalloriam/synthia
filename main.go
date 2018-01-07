package main

import (
	"time"

	"github.com/dalloriam/synthia/synthia"
)

func main() {
	sine1 := synthia.NewSine(261.6)
	sine2 := synthia.NewSine(329.6)
	sine3 := synthia.NewSine(392.0)

	speaker, err := synthia.NewSpeaker(44100, 2, 16, 9000)

	if err != nil {
		panic(err)
	}

	m := synthia.NewMixer(3)

	m.Channels[0].Input = sine1
	m.Channels[1].Input = sine2
	m.Channels[2].Input = sine3

	speaker.Input = m

	speaker.Start()
	time.Sleep(2 * time.Second)
	speaker.Stop()
}
