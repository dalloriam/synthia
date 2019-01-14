package modular

import (
	"math"

	"github.com/dalloriam/synthia/core"
)

const (
	sampleRate = 44100.0
	twoPi      = 2 * math.Pi
)

// WaveShape represents the possible shapes of an oscillator wave.
type WaveShape int

// Available shapes.
const (
	SINE WaveShape = iota
	SQUARE
	SAW
	TRIANGLE
)

// An Oscillator is a simple wave generator
type Oscillator struct {
	Frequency *core.Knob // Note frequency (in Hz).

	Sine, Square, Saw, Triangle core.Signal // From -1 to 1

	phase float64
}

// NewOscillator returns a new oscillator.
func NewOscillator() *Oscillator {
	freq := core.NewKnob(440)

	return &Oscillator{
		Frequency: freq,
		Sine:      &toneGenerator{0.0, freq, generateSine},
		Square:    &toneGenerator{0.0, freq, generateSquare},
		Saw:       &toneGenerator{0.0, freq, generateSaw},
		Triangle:  &toneGenerator{0.0, freq, generateTriangle},
	}
}

// GetOutput returns the signal corresponding to the provided wave shape.
func (o *Oscillator) GetOutput(shape WaveShape) core.Signal {
	var line core.Signal

	switch shape {
	case SINE:
		line = o.Sine
	case SQUARE:
		line = o.Square
	case TRIANGLE:
		line = o.Triangle
	case SAW:
		line = o.Saw
	default:
		line = o.Sine
	}

	return line
}

type toneGenerator struct {
	phase     float64
	frequency *core.Knob
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
	return t.tone(t.phase)

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
	return phase / twoPi
}

func generateTriangle(phase float64) float64 {
	at := phase / twoPi
	if at > 0.5 {
		at = 1.0 - at
	}
	return at*4.0 - 1.0
}
