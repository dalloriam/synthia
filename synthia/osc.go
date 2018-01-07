package synthia

import "math"

const samplesPerSecond uint32 = 44100

type Sine struct {
	radians float64
	phase   float64
}

func (s *Sine) Stream(p []float64) (int, error) {
	nbOfSamples := len(p)

	for i := 0; i < nbOfSamples; i++ {
		s.phase += s.radians
		sin := math.Sin(s.phase)
		p[i] = sin * math.MaxUint16 / 2 // TODO: Allow variable volumes
	}

	return nbOfSamples, nil
}

func NewSine(frequency float64) *Sine {
	frequencyRadiansPerSample := frequency * math.Pi / float64(samplesPerSecond)

	return &Sine{radians: frequencyRadiansPerSample, phase: 0.0}
}
