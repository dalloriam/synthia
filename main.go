package main

import (
	"github.com/dalloriam/synthia/synthia"
	"github.com/dalloriam/synthia/synthia/modular"
)

func main() {

	synth, err := synthia.NewSynth(1)
	if err != nil {
		panic(err)
	}

	sineOscillator := modular.NewOscillator(261.6, modular.SINE)

	sequencer := modular.NewSequencer([]float64{261.6, 293.7, 329.6}, 1000)

	sineOscillator.Frequency.Line = sequencer

	synth.Mixer.Channels[0].Input = sineOscillator

	select {}
}
