package modular

import (
	"math"
)

// An LFO is a Low-Frequency Oscillator. It differs from the actual oscillator in that it allows for setting a floor and a ceiling.
type LFO struct {
	osc      Oscillator
	maxValue float64
	minValue float64
	Shape    WaveShape
}

// NewLFO returns an new LFO.
func NewLFO(maxValue, minValue float64) *LFO {
	internalOsc := *NewOscillator()
	internalOsc.Frequency.SetValue(4)
	internalOsc.Volume.SetValue(math.MaxFloat64)

	return &LFO{
		osc:      internalOsc,
		maxValue: maxValue,
		minValue: minValue,
		Shape:    SINE,
	}
}

// Stream returns the current LFO phase.
func (l *LFO) Stream() float64 {

	line := l.osc.GetOutput(l.Shape)

	sample := line.Stream() + 1
	return ((sample / 2) * (l.maxValue - l.minValue)) + l.minValue
}
