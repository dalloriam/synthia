package core

// MixerChannel represents a single mixer channel
type MixerChannel struct {
	Input  Signal // From -1 to 1
	Volume *Knob  // From 0 to 1
	Pan    *Knob  // From -1 to 1
}

// NewMixerChannel initializes a mixer channel
func NewMixerChannel() *MixerChannel {
	return &MixerChannel{
		Input:  nil,
		Volume: NewKnob(0.8), // Volume initialized @ 80%,
		Pan:    NewKnob(0),   // Pan centered by default
	}
}

// Stream reads from the channel input and applies the mixer channel volume to the audio stream. It then applies panning
// if necessary.
func (c *MixerChannel) Stream(l, r []float64) {

	bufLen := len(r)

	if c.Input == nil {
		for i := 0; i < bufLen; i++ {
			r[i] = 0
			l[i] = 0
		}
	} else {

		for i := 0; i < bufLen; i++ {

			currentSample := c.Input.Stream() * c.Volume.Stream()
			pan := c.Pan.Stream()

			if pan == 0 {
				r[i] = currentSample
				l[i] = currentSample
			} else if pan > 0 {
				r[i] = (1 - pan) * currentSample
				l[i] = currentSample
			} else {
				r[i] = currentSample
				l[i] = (1 + pan) * currentSample
			}
		}
	}
}
