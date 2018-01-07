package main

import (
	"time"

	"github.com/dalloriam/synthia/synthia"
)

func main() {

	osc1 := synthia.NewOscillator(261.6, synthia.SINE)

	speaker, err := synthia.NewSpeaker(44100, 2, 16, 9000)

	if err != nil {
		panic(err)
	}

	m := synthia.NewMixer(1)

	m.Channels[0].Input = osc1

	speaker.Input = m

	speaker.Start()
	time.Sleep(2 * time.Second)
	speaker.Stop()
}
