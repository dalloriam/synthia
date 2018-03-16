package synthia

import "math"

// MixerChannel represents a single mixer channel
type MixerChannel struct {
	Input  Signal
	Volume *Knob // From 0 to Float64Max
	Pan    *Knob // From -1 to 1
}

// NewMixerChannel initializes a mixer channel
func NewMixerChannel() *MixerChannel {
	return &MixerChannel{
		Input:  nil,
		Volume: NewKnob(204), // Volume initialized @ 80%,
		Pan:    NewKnob(0),   // Pan centered by default
	}
}

// Stream reads from the channel input and applies the mixer channel volume to the audio stream
func (c *MixerChannel) Stream(l, r []float64) {

	bufLen := len(r)

	if c.Input == nil {
		for i := 0; i < bufLen; i++ {
			r[i] = 0
			l[i] = 0
		}
	} else {

		volBuf := make([]float64, bufLen)
		c.Volume.Stream(volBuf)

		panBuf := make([]float64, bufLen)
		c.Pan.Stream(panBuf)

		p := make([]float64, bufLen)

		c.Input.Stream(p)

		for i := 0; i < bufLen; i++ {
			volFactor := volBuf[i] / float64(math.MaxUint8)

			currentSample := p[i] * volFactor

			if panBuf[i] == 0 {
				r[i] = currentSample
				l[i] = currentSample
			} else if panBuf[i] > 0 {
				r[i] = (1 - panBuf[i]) * currentSample
				l[i] = currentSample
			} else {
				r[i] = currentSample
				l[i] = (1 + panBuf[i]) * currentSample
			}
		}
	}
}
