package modular

import (
	"math"

	"github.com/dalloriam/synthia"
)

// WaveShape represents a wave in the internal oscillator wavetable
type WaveShape int

const (
	SINE = iota
	SQUARE
)

const sampleRate = 44100.0

// An Oscillator is a simple wave generator
type Oscillator struct {
	Frequency *synthia.Knob

	Shape  WaveShape
	Volume *synthia.Knob

	phase float64
}

// NewOscillator returns a new oscillator.
func NewOscillator(freq float64, shape WaveShape) *Oscillator {
	return &Oscillator{
		Frequency: synthia.NewKnob(freq),
		Shape:     shape,
		Volume:    synthia.NewKnob(math.MaxFloat64),
	}
}

func (o *Oscillator) incrementPhase(freq float64) {
	if o.phase > 2*math.Pi {
		o.phase = o.phase - (2 * math.Pi)
	}
	o.phase += freq * math.Pi / sampleRate
}

func (o *Oscillator) sine(p []float64) {
	nbOfSamples := len(p)

	volBuf := make([]float64, len(p))
	o.Volume.Stream(volBuf)

	phaseBuf := make([]float64, len(p))
	o.Frequency.Stream(phaseBuf)

	for i := 0; i < nbOfSamples; i++ {
		volFactor := volBuf[i] / math.MaxFloat64
		o.incrementPhase(phaseBuf[i])
		sin := math.Sin(o.phase)

		p[i] = sin * volFactor * math.MaxUint16 / 2
	}

}

func (o *Oscillator) square(p []float64) {
	nbOfSamples := len(p)

	var wv float64

	volBuf := make([]float64, len(p))
	o.Volume.Stream(volBuf)

	phaseBuf := make([]float64, len(p))
	o.Frequency.Stream(phaseBuf)

	for i := 0; i < nbOfSamples; i++ {
		o.incrementPhase(phaseBuf[i])
		if math.Sin(o.phase) > 0 {
			wv = math.MaxUint16
		} else {
			wv = 0
		}
		p[i] = wv * (volBuf[i] / math.MaxFloat64)
	}
}

// Stream writes the current phase to the buffer.
func (o *Oscillator) Stream(p []float64) {
	switch o.Shape {
	case SINE:
		o.sine(p)
	case SQUARE:
		o.square(p)
	}
}
