package modular

import (
	"math"

	"github.com/dalloriam/synthia"
)

// An LFO is a Low-Frequency Oscillator. It differs from the actual oscillator in that it allows for setting a floor and a ceiling.
type LFO struct {
	Oscillator
	maxValue float64
	minValue float64
}

// NewLFO returns an new LFO.
func NewLFO(maxValue, minValue float64) *LFO {
	return &LFO{
		Oscillator: Oscillator{
			Frequency: synthia.NewKnob(4),
			Shape:     SINE,
			Volume:    synthia.NewKnob(math.MaxFloat64),
		},
		maxValue: maxValue,
		minValue: minValue,
	}
}

// Stream writes the current LFO phase to the audio buffer.
func (l *LFO) Stream(p []float64) {

	l.Oscillator.Stream(p)

	for i := 0; i < len(p); i++ {
		p[i] = ((p[i] / math.MaxUint16) * (l.maxValue - l.minValue)) + l.minValue
	}
}
