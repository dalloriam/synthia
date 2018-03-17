package modular

import (
	"time"

	"github.com/dalloriam/synthia"
)

type Clock struct {
	startTime       time.Time
	Tempo           *synthia.Knob // In bpm
	TicksPerQuarter int
}

func NewClock() *Clock {
	return &Clock{
		startTime:       time.Now(),
		Tempo:           synthia.NewKnob(60),
		TicksPerQuarter: 24,
	}
}

func (c *Clock) Stream() float64 {

	ticksPerBeat := c.TicksPerQuarter * 4

	clockTicksPerSecond := c.Tempo.Stream() * float64(ticksPerBeat) / 60.0

	nbOfTicks := int(time.Since(c.startTime).Seconds() * clockTicksPerSecond)

	return float64(nbOfTicks)
}
