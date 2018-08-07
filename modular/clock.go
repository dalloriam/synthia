package modular

import (
	"time"

	"github.com/dalloriam/synthia/core"
)

const clockTicksPerBeat = 96.0

// A Clock synchronizes all modules connected to it by emitting a steady stream of pulses (96/beat by default).
type Clock struct {
	startTime time.Time
	Tempo     *core.Knob // In beats per minute
}

// NewClock returns a clock initialized with a tempo of 60bpm and a tick rate of 96 ticks / beat.
func NewClock() *Clock {
	return &Clock{
		startTime: time.Now(),
		Tempo:     core.NewKnob(60),
	}
}

// Stream reads the current value of the clock
func (c *Clock) Stream() float64 {

	tempo := c.Tempo.Stream()

	clockTicksPerSecond := tempo * float64(clockTicksPerBeat) / 60.0

	nbOfTicks := int(time.Since(c.startTime).Seconds() * clockTicksPerSecond)

	return float64(nbOfTicks)
}
