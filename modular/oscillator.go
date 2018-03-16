package modular

import (
	"math"

	"github.com/Dalloriam/synthia"
)

// WaveShape represents a wave in the internal oscillator wavetable
type WaveShape int

const (
	SINE = iota
	SQUARE
	SAWTOOTH
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

func (o *Oscillator) getToneGenerator() func() float64 {
	var wave func() float64

	switch o.Shape {
	case SQUARE:
		wave = o.square

	case SINE:
		wave = o.sine
	case SAWTOOTH:
		wave = o.sawtooth
	}

	return wave
}

func (o *Oscillator) incrementPhase(freq float64) {
	o.phase += freq * 2 * math.Pi / sampleRate
	if o.phase > 2*math.Pi {
		o.phase -= 2 * math.Pi
	}
}

func (o *Oscillator) sine() float64 {
	return math.Sin(o.phase)
}

func (o *Oscillator) square() float64 {
	if math.Sin(o.phase) > 0 {
		return 1
	} else {
		return -1
	}
}

func (o *Oscillator) sawtooth() float64 {
	p := o.phase / (2 * math.Pi)
	return (2 * p) - 1
}

// Stream writes the current phase to the buffer.
func (o *Oscillator) Stream(p []float64) {
	toneGenerator := o.getToneGenerator()

	nbOfSamples := len(p)

	volBuf := make([]float64, len(p))
	o.Volume.Stream(volBuf)

	freqBuf := make([]float64, len(p))
	o.Frequency.Stream(freqBuf)

	for i := 0; i < nbOfSamples; i++ {
		o.incrementPhase(freqBuf[i])
		p[i] = toneGenerator() * (volBuf[i] / math.MaxFloat64) * math.MaxUint16 / 2
	}
}
