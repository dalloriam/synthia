package modular

import "github.com/dalloriam/synthia/core"

// A Sequencer loops through a frequency sequence at a pace dictated by the clock
// and outputs the corresponding frequencies.
type Sequencer struct {
	Clock        core.Signal
	Sequence     []float64
	BeatsPerStep *core.Knob
	Trigger      *Trigger

	lastStepIdx int
}

// NewSequencer returns a sequencer instance.
func NewSequencer(sequence []float64) *Sequencer {
	return &Sequencer{
		Sequence:     sequence,
		BeatsPerStep: core.NewKnob(0.5),
		Trigger:      NewTrigger(),
		lastStepIdx:  -1,
	}
}

// Stream returns the current sequence frequency.
func (s *Sequencer) Stream() float64 {
	// TODO: Support clock-free looping (time-based)

	ticksPerStep := float64(clockTicksPerBeat) * s.BeatsPerStep.Stream()

	currentClock := s.Clock.Stream()

	stepIdx := int(currentClock/ticksPerStep) % len(s.Sequence)

	if stepIdx != s.lastStepIdx {
		s.lastStepIdx = stepIdx
		s.Trigger.ShouldTrigger = true
	}

	return s.Sequence[stepIdx]
}
