package synthia

const sampleRate = 44100
const audioChannelCount = 2
const bitsPerSample = 16
const bufferSize = 8000 // TODO: Adjust dynamically

type Synthia struct {
	sampleRate int
	speaker    *Speaker
	Mixer      *Mixer
}

func NewSynth(channelCount int) (*Synthia, error) {
	m := NewMixer(channelCount)
	spk, err := NewSpeaker(audioChannelCount, bitsPerSample, bufferSize, sampleRate)

	if err != nil {
		return nil, err
	}

	spk.Input = m
	spk.Start()
	return &Synthia{sampleRate: sampleRate, Mixer: m, speaker: spk}, nil
}
