package core

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

	return &Mixer{Channels: chans}
}

// Stream mixes all the mixer channels in a single stereo stream
func (m *Mixer) Stream(l, r []float64) {
	// TODO: Improve

	bufSize := len(r)

	chanCnt := len(m.Channels)

	lbufs := make([][]float64, chanCnt)
	rbufs := make([][]float64, chanCnt)

	rmaxes := make([]float64, chanCnt)
	lmaxes := make([]float64, chanCnt)

	for i := 0; i < chanCnt; i++ {
		lbufs[i] = make([]float64, bufSize)
		rbufs[i] = make([]float64, bufSize)
		m.Channels[i].Stream(lbufs[i], rbufs[i])

		for j := 0; j < len(rbufs[i]); j++ {
			if rbufs[i][j] > rmaxes[i] {
				rmaxes[i] = rbufs[i][j]
			}
			if lbufs[i][j] > lmaxes[i] {
				lmaxes[i] = lbufs[i][j]
			}
		}
	}

	rSum := 0.0
	lSum := 0.0

	for i := 0; i < chanCnt; i++ {
		rSum += rmaxes[i]
		lSum += lmaxes[i]
	}

	for i := 0; i < chanCnt; i++ {
		rmaxes[i] = rmaxes[i] / rSum
		lmaxes[i] = lmaxes[i] / lSum
	}

	for j := 0; j < len(m.Channels); j++ {

		for i := 0; i < len(rbufs[j]); i++ {
			l[i] = l[i] + (lbufs[j][i] * lmaxes[j])
			r[i] = r[i] + (rbufs[j][i] * rmaxes[j])
		}
	}
}
