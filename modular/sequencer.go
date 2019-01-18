package modular

import (
	"github.com/dalloriam/synthia/core"
	"github.com/dalloriam/synthia/core/constants"
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

	currentStep int // Current step stores the current step in the sequence.
	tickCounter int // Pulse counter counts the number of clock pulses since the last step.
}

// NewSequencer returns a sequencer instance.
func NewSequencer() *Sequencer {
	return &Sequencer{
		PitchSequence: []float64{130.81, 146.83, 164.1, 174.61, 196, 220, 246.94, 261.63},
		GateSequence:  []float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5},
		BeatsPerStep:  core.NewKnob(0.5),
		Gate:          NewGate(),
		Trigger:       NewTrigger(),

		currentStep: 0,
		tickCounter: 0,
	}
}

func (s *Sequencer) incrementStep() {
	if s.currentStep+1 >= len(s.PitchSequence) {
		s.currentStep = 0
	} else {
		s.currentStep++
	}
}

// Stream returns the current sequence frequency.
func (s *Sequencer) Stream() float64 {
	ticksPerStep := int(constants.ClockTicksPerBeat * s.BeatsPerStep.Stream())
	currentClock := s.Clock.Stream()

	if currentClock == 1 {
		s.tickCounter++
	}

	if s.tickCounter >= ticksPerStep {
		s.tickCounter = 0
		s.Trigger.ShouldTrigger = true
		s.incrementStep()
	}

	percentageToNextStep := float64(s.tickCounter) / float64(ticksPerStep)
	gateSequenceValue := s.GateSequence[s.currentStep]

	if percentageToNextStep > gateSequenceValue {
		s.Gate.Opened = false
	} else {
		s.Gate.Opened = true
	}

	return s.PitchSequence[s.currentStep]
}
