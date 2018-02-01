package synthia

// A Mixer is a collection of mixer channels
type Mixer struct {
	Channels []*MixerChannel
}

// NewMixer returns a new mixer with all of its channels already initialized
func NewMixer(nbOfChannels int) *Mixer {

	chans := make([]*MixerChannel, nbOfChannels)
	for i := 0; i < nbOfChannels; i++ {
		chans[i] = NewMixerChannel()
	}

	return &Mixer{chans}
}

// Stream mixes all the mixer channels in a single audio stream
func (m *Mixer) Stream(p []float64) {

	chanCnt := len(m.Channels)

	bufs := make([][]float64, chanCnt)

	maxes := make([]float64, chanCnt)

	for i := 0; i < chanCnt; i++ {
		bufs[i] = make([]float64, len(p))
		m.Channels[i].Stream(bufs[i])

		for j := 0; j < len(bufs[i]); j++ {
			if bufs[i][j] > maxes[i] {
				maxes[i] = bufs[i][j]
			}
		}
	}

	maxSum := 0.0
	for _, mx := range maxes {
		maxSum += mx
	}

	for i := 0; i < chanCnt; i++ {
		maxes[i] = maxes[i] / maxSum
	}

	for j := 0; j < len(m.Channels); j++ {

		for i := 0; i < len(bufs[j]); i++ {
			p[i] = p[i] + (bufs[j][i] * maxes[j])
		}
	}
}
