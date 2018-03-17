package modular

import (
	"github.com/dalloriam/synthia"
)

// A Sequencer loops through a note sequence and outputs the corresponding frequencies to a stream
type Sequencer struct {
	Clock       synthia.Signal
	Sequence    []float64
	lastClock   float64
	currentStep int
}

// NewSequencer returns a sequencer instance.
func NewSequencer(sequence []float64) *Sequencer {
	return &Sequencer{
		Sequence:    sequence,
		lastClock:   -24,
		currentStep: -1,
	}
}

// Stream writes the current sequence frequency to the buffer
func (s *Sequencer) Stream() float64 {

	// Play quarternotes by default
	currentClock := s.Clock.Stream()

	if currentClock > s.lastClock && (currentClock-s.lastClock) >= 24 {
		s.lastClock = currentClock
		s.currentStep++
	}

	// TODO: Watch for overflow on s.currentStep
	v := s.currentStep % len(s.Sequence)
	return s.Sequence[v]
}
