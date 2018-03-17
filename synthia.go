package synthia

// Synthia is the core synthesizer struct
type Synthia struct {
	Speaker *Speaker
	Mixer   *Mixer
}

type audioBackend interface {
	Start(callback func(in []float32, out [][]float32)) error
	FrameSize() int
}

// NewSynth returns a new synthesizer with an already-initialized mixer
func NewSynth(channelCount, bufferSize int, output audioBackend) *Synthia {
	m := NewMixer(channelCount)
	spk := NewSpeaker(bufferSize, output.FrameSize())
	spk.Input = m

	err := output.Start(spk.ProcessBuffer)
	if err != nil {
		panic(err)
	}

	return &Synthia{Mixer: m, Speaker: spk}
}
