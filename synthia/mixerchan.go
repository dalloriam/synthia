package synthia

import "math"

type MixerChannel struct {
	Input  AudioStream
	Volume byte
	Output AudioStream
}

func NewMixerChannel() *MixerChannel {
	return &MixerChannel{
		Input:  nil,
		Volume: 204, // Volume initialized @ 80%
	}
}

func (c *MixerChannel) Stream(p []float64) {
	if c.Input == nil {
		for i := 0; i < len(p); i++ {
			p[i] = 0
		}
	} else {
		c.Input.Stream(p)
		volFactor := float64(c.Volume) / float64(math.MaxUint8)
		for i := 0; i < len(p); i++ {
			p[i] = p[i] * volFactor
		}
	}
}
