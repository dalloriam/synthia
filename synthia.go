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
func NewSynth(channelCount int) (*Synthia, error) {
	m := NewMixer(channelCount)
	spk, err := NewSpeaker(audioChannelCount, bitsPerSample, bufferSize, sampleRate)

	if err != nil {
		return nil, err
	}

	spk.Input = m
	spk.Start()
	return &Synthia{Mixer: m, speaker: spk}, nil
}
