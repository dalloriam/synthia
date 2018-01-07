package main

import (
	"time"

	"github.com/dalloriam/synthia/synthia"
)

func main() {

	osc1 := synthia.NewOscillator(261.6, synthia.SINE)
	osc2 := synthia.NewOscillator(440.0, synthia.SINE)

	speaker, err := synthia.NewSpeaker(2, 16, 9000)

	if err != nil {
		panic(err)
	}

	m := synthia.NewMixer(2)

	m.Channels[0].Input = osc1
	m.Channels[1].Input = osc2

	speaker.Input = m

	speaker.Start()
	time.Sleep(2 * time.Second)
	speaker.Stop()
}
