package modular

// A Sequencer loops through a frequency sequence at a pace dictated by the clock
// and outputs the corresponding frequencies.
type Sequencer struct {
	Clock        *Clock
	Sequence     []float64
	BeatsPerStep float64

	lastClock float64
}

// NewSequencer returns a sequencer instance.
func NewSequencer(sequence []float64) *Sequencer {
	return &Sequencer{
		Sequence:     sequence,
		lastClock:    -24,
		BeatsPerStep: 0.5,
	}
}

// Stream returns the current sequence frequency.
func (s *Sequencer) Stream() float64 {
	// TODO: Support clock-free looping (time-based)

	ticksPerStep := float64(s.Clock.TicksPerBeat) * s.BeatsPerStep

	currentClock := s.Clock.Stream()

	stepIdx := int(currentClock/ticksPerStep) % len(s.Sequence)

	return s.Sequence[stepIdx]
}
