package modular

import (
	"time"

	"github.com/dalloriam/synthia/core"
)

// A Clock synchronizes all modules connected to it by emitting a steady stream of pulses (96/beat by default).
type Clock struct {
	startTime    time.Time
	Tempo        *core.Knob // In beats per minute
	TicksPerBeat int
}

// NewClock returns a clock initialized with a tempo of 60bpm and a tick rate of 96 ticks / beat.
func NewClock() *Clock {
	return &Clock{
		startTime:    time.Now(),
		Tempo:        core.NewKnob(60),
		TicksPerBeat: 96,
	}
}

// Stream reads the current value of the clock
func (c *Clock) Stream() float64 {

	tempo := c.Tempo.Stream()

	clockTicksPerSecond := tempo * float64(c.TicksPerBeat) / 60.0

	nbOfTicks := int(time.Since(c.startTime).Seconds() * clockTicksPerSecond)

	return float64(nbOfTicks)
}
