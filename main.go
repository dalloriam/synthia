package main

import (
	"github.com/dalloriam/synthia/synthia"
	"github.com/dalloriam/synthia/synthia/modular"
)

func main() {

	synth, err := synthia.NewSynth(2)
	if err != nil {
		panic(err)
	}

	sineOscillator := modular.NewOscillator(261.6, modular.SINE)
	synth.Mixer.Channels[0].Input = sineOscillator

	sineOscillator2 := modular.NewOscillator(440.0, modular.SINE)
	synth.Mixer.Channels[1].Input = sineOscillator2

	select {}
}
