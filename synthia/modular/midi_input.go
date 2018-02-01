package modular

type constantStreamer struct {
	value float64
}

func (c *constantStreamer) Stream(p []float64) {
	for i := 0; i < len(p); i++ {
		p[i] = c.value
	}
}

type MIDIInput struct {
	VolumeControl    *constantStreamer
	FrequencyControl *constantStreamer
}

func NewMIDIInput() *MIDIInput {
	return &MIDIInput{
		VolumeControl:    &constantStreamer{value: 0},
		FrequencyControl: &constantStreamer{value: 440},
	}
}

func (m *MIDIInput) UpdateVolume(volume float64) {
	m.VolumeControl.value = volume
}

func (m *MIDIInput) UpdateFrequency(frequency float64) {
	m.FrequencyControl.value = frequency
}
