package synthia

const sampleRate = 44100
const audioChannelCount = 2
const bitsPerSample = 16
const bufferSize = 8000 // TODO: Adjust dynamically

// Synthia is the core synthesizer struct
type Synthia struct {
	speaker *Speaker
	Mixer   *Mixer
}

// NewSynth returns a new synthesizer with an already-initialized mixer
func NewSynth(channelCount int, output StreamOutput) *Synthia {
	m := NewMixer(channelCount)
	spk := NewSpeaker(output, bufferSize)

	spk.Input = m
	spk.Start()
	return &Synthia{Mixer: m, speaker: spk}
}
