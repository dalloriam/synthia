package synthia

type Mixer struct {
	Channels []*MixerChannel
}

func NewMixer(nbOfChannels int) *Mixer {

	chans := make([]*MixerChannel, nbOfChannels)
	for i := 0; i < nbOfChannels; i++ {
		chans[i] = NewMixerChannel()
	}

	return &Mixer{chans}
}

func (m *Mixer) Stream(p []float64) (int, error) {

	realChanCount := 0

	for _, s := range m.Channels {
		channelNull := true

		buf := make([]float64, len(p))
		s.Stream(buf)

		for i := 0; i < len(p); i++ {
			x := p[i] + buf[i]
			p[i] = x
			if buf[i] != 0 {
				channelNull = false
			}
		}

		if !channelNull {
			realChanCount++
		}
	}

	if realChanCount > 1 {
		for i := 0; i < len(p); i++ {
			p[i] = p[i] / float64(realChanCount)
		}
	}

	return len(p), nil
}
