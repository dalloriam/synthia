package main

import (
	"math"

	"github.com/dalloriam/synthia/synthia"
)

func main() {

	osc1 := synthia.NewOscillator(261.6, synthia.SINE)

	speaker, err := synthia.NewSpeaker(2, 16, 8000)

	if err != nil {
		panic(err)
	}

	m := synthia.NewMixer(2)

	m.Channels[0].Input = osc1

	speaker.Input = m

	lfolfo := synthia.NewLFO(8.0, 1.0)
	lfolfo.Frequency.SetValue(0.125)

	lfo := synthia.NewLFO(math.MaxFloat64, 0)
	lfo.Frequency.Line = lfolfo

	sequencer := synthia.NewSequencer([]float64{261.6, 293.7, 329.6}, 1000)

	osc1.Frequency.Line = sequencer
	osc1.Volume.Line = lfo

	speaker.Start()

	select {}
}
