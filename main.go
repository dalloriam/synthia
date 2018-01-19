package main

import "github.com/dalloriam/synthia/synthia"

func main() {

	osc1 := synthia.NewOscillator(261.6, synthia.SINE)

	speaker, err := synthia.NewSpeaker(2, 16, 10000)

	if err != nil {
		panic(err)
	}

	m := synthia.NewMixer(2)

	m.Channels[0].Input = osc1

	speaker.Input = m

	sequencer := synthia.NewSequencer([]float64{261.6, 293.7, 329.6}, 300)

	osc1.Frequency.Line = sequencer

	speaker.Start()

	select {}
}
