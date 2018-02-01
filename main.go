package main

import (
	"math"
	"time"

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

	inpt := modular.NewMIDIInput()

	sineOscillator.Frequency.Line = inpt.FrequencyControl
	sineOscillator.Volume.Line = inpt.VolumeControl

	inpt.UpdateFrequency(400)
	inpt.UpdateVolume(math.MaxFloat64)

	time.Sleep(2 * time.Second)

	inpt.UpdateVolume(0)

	select {}
}
