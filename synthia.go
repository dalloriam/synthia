package synthia

// Synthia is the core synthesizer struct
type Synthia struct {
	speaker *Speaker
	Mixer   *Mixer
}

// NewSynth returns a new synthesizer with an already-initialized mixer
func NewSynth(channelCount, bufferSize int, output StreamOutput) *Synthia {
	m := NewMixer(channelCount)
	spk := NewSpeaker(output, bufferSize)

	spk.InputR = m
	spk.Start()
	return &Synthia{Mixer: m, speaker: spk}
}
