package modular

import (
	"math"

	"github.com/Dalloriam/synthia"
)

// WaveShape represents a wave in the internal oscillator wavetable
type WaveShape int

const sampleRate = 44100.0

// An Oscillator is a simple wave generator
type Oscillator struct {
	Frequency *synthia.Knob
	Volume    *synthia.Knob

	Sine   synthia.Signal
	Square synthia.Signal
	Saw    synthia.Signal

	phase float64
}

// NewOscillator returns a new oscillator.
func NewOscillator() *Oscillator {
	vol := synthia.NewKnob(math.MaxFloat64)
	freq := synthia.NewKnob(440)

	return &Oscillator{
		Frequency: freq,
		Volume:    vol,
		Sine:      &toneGenerator{0.0, vol, freq, generateSine},
		Square:    &toneGenerator{0.0, vol, freq, generateSquare},
		Saw:       &toneGenerator{0.0, vol, freq, generateSaw},
	}
}

type toneGenerator struct {
	phase     float64
	volume    *synthia.Knob
	frequency *synthia.Knob
	tone      func(phase float64) float64
}

func (t *toneGenerator) incrementPhase(freq float64) {
	t.phase += freq * 2 * math.Pi / sampleRate
	if t.phase > 2*math.Pi {
		t.phase -= 2 * math.Pi
	}
}

func (t *toneGenerator) Stream(p []float64) {
	nbOfSamples := len(p)

	volBuf := make([]float64, len(p))
	t.volume.Stream(volBuf)

	freqBuf := make([]float64, len(p))
	t.frequency.Stream(freqBuf)

	for i := 0; i < nbOfSamples; i++ {
		t.incrementPhase(freqBuf[i])
		p[i] = t.tone(t.phase) * (volBuf[i] / math.MaxFloat64) * math.MaxUint16 / 2
	}
}

func generateSine(phase float64) float64 {
	return math.Sin(phase)
}

func generateSquare(phase float64) float64 {
	if math.Sin(phase) > 0 {
		return 1
	}
	return -1
}

func generateSaw(phase float64) float64 {
	p := phase / (2 * math.Pi)
	return (2 * p) - 1
}
