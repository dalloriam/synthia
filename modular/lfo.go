package modular

import (
	"math"
)

// An LFO is a Low-Frequency Oscillator. It differs from the actual oscillator in that it allows for setting a floor and a ceiling.
type LFO struct {
	Oscillator
	maxValue float64
	minValue float64
}

// NewLFO returns an new LFO.
func NewLFO(maxValue, minValue float64) *LFO {
	internalOsc := *NewOscillator()
	internalOsc.Frequency.SetValue(4)
	internalOsc.Volume.SetValue(math.MaxFloat64)

	return &LFO{
		Oscillator: internalOsc,
		maxValue:   maxValue,
		minValue:   minValue,
	}
}

// Stream writes the current LFO phase to the audio buffer.
func (l *LFO) Stream() float64 {
	sample := l.Oscillator.Sine.Stream() + 1
	return ((sample / 2) * (l.maxValue - l.minValue)) + l.minValue
}
