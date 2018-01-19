package modular

import "time"

type Sequencer struct {
	Sequence      []float64
	StepFrequency *Knob
	startTime     time.Time
}

func NewSequencer(sequence []float64, stepDelay float64) *Sequencer {
	return &Sequencer{
		Sequence:      sequence,
		StepFrequency: NewKnob(stepDelay),
		startTime:     time.Now(),
	}
}

func (s *Sequencer) Stream(p []float64) (int, error) {

	stepBuf := make([]float64, len(p))
	s.StepFrequency.Stream(stepBuf)

	for i := 0; i < len(p); i++ {
		numOfSteps := int((time.Since(s.startTime).Seconds() * 1000.0) / float64(stepBuf[i]))

		v := numOfSteps % len(s.Sequence)

		if numOfSteps > 1000 && v == 0 {
			s.startTime = time.Now()
		}

		p[i] = s.Sequence[v]
	}
	return len(p), nil
}
