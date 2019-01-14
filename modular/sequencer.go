package modular

import (
	"github.com/dalloriam/synthia/core"
)

// A Sequencer loops through a frequency sequence at a pace dictated by the clock
// and outputs the corresponding frequencies.
type Sequencer struct {
	Clock        core.Signal
	BeatsPerStep *core.Knob

	Gate    *Gate
	Trigger *Trigger

	PitchSequence []float64
	GateSequence  []float64

	lastStepIdx int
}

// NewSequencer returns a sequencer instance.
func NewSequencer() *Sequencer {
	return &Sequencer{
		PitchSequence: []float64{130.81, 146.83, 164.1, 174.61, 196, 220, 246.94, 261.63},
		GateSequence:  []float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5},
		BeatsPerStep:  core.NewKnob(0.5),
		Gate:          NewGate(),
		Trigger:       NewTrigger(),
		lastStepIdx:   -1,
	}
}

// Stream returns the current sequence frequency.
func (s *Sequencer) Stream() float64 {
	ticksPerStep := float64(clockTicksPerBeat) * s.BeatsPerStep.Stream()

	currentClock := s.Clock.Stream()

	stepIdx := int(currentClock/ticksPerStep) % len(s.PitchSequence)

	percentageToNextStep := float64(int(currentClock)%int(ticksPerStep)) / ticksPerStep
	gateSequenceValue := s.GateSequence[stepIdx]
	if percentageToNextStep > gateSequenceValue {
		s.Gate.Opened = false
	} else {
		s.Gate.Opened = true
	}

	if stepIdx != s.lastStepIdx {
		s.lastStepIdx = stepIdx
		s.Trigger.ShouldTrigger = true
	}

	return s.PitchSequence[stepIdx]
}
