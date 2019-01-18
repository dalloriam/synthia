package modular

import (
	"github.com/dalloriam/synthia/core"
	"github.com/dalloriam/synthia/core/constants"
)

// A Clock synchronizes all modules connected to it by emitting a steady stream of ticks (96 ticks per beat).
type Clock struct {
	Tempo    *core.Knob // In beats per minute.
	position int
}

// NewClock returns a clock initialized with a tempo of 60bpm and a tick rate of 96 ticks / beat.
func NewClock() *Clock {
	return &Clock{
		Tempo:    core.NewKnob(60),
		position: 0,
	}
}

// Stream reads the current value of the clock
func (c *Clock) Stream() float64 {
	tempo := c.Tempo.Stream()

	ticksPerSecond := int(tempo * constants.ClockTicksPerBeat / 60)
	samplesPerTick := constants.SampleRate / ticksPerSecond

	if c.position != 0 && c.position%samplesPerTick == 0 {
		// Reset position and emit pulse
		c.position = 0
		return 1
	}
	c.position++
	return 0
}
