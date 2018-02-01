package synthia

import "math"

type MixerChannel struct {
	Input  AudioStream
	Volume *Knob
}

func NewMixerChannel() *MixerChannel {
	return &MixerChannel{
		Input:  nil,
		Volume: NewKnob(204), // Volume initialized @ 80%
	}
}

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
			volFactor := volBuf[i] / float64(math.MaxUint8)
			p[i] = p[i] * volFactor
		}
	}
}
