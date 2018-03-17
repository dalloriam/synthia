package modular

import (
	"math"

	"github.com/dalloriam/synthia"
)

// WaveShape represents a wave in the internal oscillator wavetable
type WaveShape int

const sampleRate = 44100.0
const twoPi = 2 * math.Pi

// An Oscillator is a simple wave generator
type Oscillator struct {
	Frequency *synthia.Knob
	Volume    *synthia.Knob

	Sine     synthia.Signal
	Square   synthia.Signal
	Saw      synthia.Signal
	Triangle synthia.Signal

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
		Triangle:  &toneGenerator{0.0, vol, freq, generateTriangle},
	}
}

type toneGenerator struct {
	phase     float64
	volume    *synthia.Knob
	frequency *synthia.Knob
	tone      func(phase float64) float64
}

func (t *toneGenerator) incrementPhase(freq float64) {
	t.phase += freq * twoPi / sampleRate
	if t.phase > twoPi {
		t.phase -= twoPi
	}
}

func (t *toneGenerator) Stream() float64 {

	t.incrementPhase(t.frequency.Stream())
	return t.tone(t.phase) * (t.volume.Stream() / math.MaxFloat64) * math.MaxUint16 / 2

}

func generateSine(phase float64) float64 {
	return sin(phase)
}

func generateSquare(phase float64) float64 {
	if sin(phase) > 0 {
		return 1
	}
	return -1
}

func generateSaw(phase float64) float64 {
	p := phase / twoPi
	return (2 * p) - 1
}

func generateTriangle(phase float64) float64 {
	at := phase / twoPi
	if at > 0.5 {
		at = 1.0 - at
	}
	return at*4.0 - 1.0
}
