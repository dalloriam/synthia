package modular

// A Sequencer loops through a note sequence and outputs the corresponding frequencies to a stream
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

// Stream writes the current sequence frequency to the buffer
func (s *Sequencer) Stream() float64 {

	ticksPerStep := float64(s.Clock.TicksPerBeat) * s.BeatsPerStep

	currentClock := s.Clock.Stream()

	stepIdx := int(currentClock/ticksPerStep) % len(s.Sequence)

	return s.Sequence[stepIdx]
}
