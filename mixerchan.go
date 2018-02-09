package synthia

import "math"

// MixerChannel represents a single mixer channel
type MixerChannel struct {
	Input  AudioStream
	Volume *Knob
}

// NewMixerChannel initializes a mixer channel
func NewMixerChannel() *MixerChannel {
	return &MixerChannel{
		Input:  nil,
		Volume: NewKnob(204), // Volume initialized @ 80%
	}
}

// Stream reads from the channel input and applies the mixer channel volume to the audio stream
func (c *MixerChannel) Stream(p []float64) {
	volBuf := make([]float64, len(p))
	c.Volume.Stream(volBuf)

	if c.Input == nil {
		for i := 0; i < len(p); i++ {
			p[i] = 0
		}
	} else {
		c.Input.Stream(p)
		for i := 0; i < len(p); i++ {
			volFactor := volBuf[i] / float64(math.MaxFloat64)
			p[i] = p[i] * volFactor
		}
	}
}
