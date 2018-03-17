package modular

import (
	"time"

	"fmt"

	"github.com/dalloriam/synthia"
)

type Clock struct {
	startTime    time.Time
	Tempo        *synthia.Knob // In bpm
	TicksPerBeat int
}

func NewClock() *Clock {
	return &Clock{
		startTime:    time.Now(),
		Tempo:        synthia.NewKnob(60),
		TicksPerBeat: 96,
	}
}

func (c *Clock) Stream() float64 {

	tempo := c.Tempo.Stream()

	clockTicksPerSecond := tempo * float64(c.TicksPerBeat) / 60.0

	nbOfTicks := int(time.Since(c.startTime).Seconds() * clockTicksPerSecond)

	v := nbOfTicks % c.TicksPerBeat

	if nbOfTicks > 1000 && v == 0 {
		c.startTime = time.Now()
	}

	fmt.Println(nbOfTicks)

	return float64(nbOfTicks)
}
