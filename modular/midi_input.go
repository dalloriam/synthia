package modular

type constantStreamer struct {
	value float64
}

func (c *constantStreamer) Stream(p []float64) {
	for i := 0; i < len(p); i++ {
		p[i] = c.value
	}
}

// A MIDIInput allows for live control of the synthesizer
type MIDIInput struct {
	VolumeControl    *constantStreamer
	FrequencyControl *constantStreamer
}

// NewMIDIInput returns a midi input
func NewMIDIInput() *MIDIInput {
	return &MIDIInput{
		VolumeControl:    &constantStreamer{value: 0},
		FrequencyControl: &constantStreamer{value: 440},
	}
}

// UpdateVolume returns the current oscillator volume
func (m *MIDIInput) UpdateVolume(volume float64) {
	m.VolumeControl.value = volume
}

// UpdateFrequency returns the current oscillator frequency
func (m *MIDIInput) UpdateFrequency(frequency float64) {
	m.FrequencyControl.value = frequency
}
