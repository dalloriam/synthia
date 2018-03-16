package synthia

import "io"

// Synthia is the core synthesizer struct
type Synthia struct {
	speaker *Speaker
	Mixer   *Mixer
}

// NewSynth returns a new synthesizer with an already-initialized mixer
func NewSynth(channelCount, bufferSize int, output io.Writer) *Synthia {
	m := NewMixer(channelCount)
	spk := NewSpeaker(output, bufferSize)

	spk.Input = m
	spk.Start()
	return &Synthia{Mixer: m, speaker: spk}
}
