package modular

import "math"

type LFO struct {
	Oscillator
	maxValue float64
	minValue float64
}

func NewLFO(maxValue, minValue float64) *LFO {
	return &LFO{
		Oscillator: Oscillator{
			Frequency: NewKnob(4),
			Shape:     SINE,
			Volume:    NewKnob(math.MaxFloat64),
		},
		maxValue: maxValue,
		minValue: minValue,
	}
}

func (l *LFO) Stream(p []float64) {

	l.Oscillator.Stream(p)

	for i := 0; i < len(p); i++ {
		p[i] = ((p[i] / math.MaxUint16) * (l.maxValue - l.minValue)) + l.minValue
	}
}
