package modular

import "github.com/dalloriam/synthia/core"

// The Modulator applies a modulation range to a signal.
// Returns InitialAmount + (Attenuator * ModulationRange).
type Modulator struct {
	Attenuator      *core.Knob // From 0 to 1.
	InitialAmount   *core.Knob
	ModulationRange *core.Knob
}

// NewModulator returns a new modulation source.
func NewModulator(initialAmount float64) *Modulator {
	return &Modulator{
		Attenuator:      core.NewKnob(0), // No modulation by default.
		InitialAmount:   core.NewKnob(initialAmount),
		ModulationRange: core.NewKnob(0), // No modulation range by default
	}
}

func (m *Modulator) Stream() float64 {
	return m.InitialAmount.Stream() + (m.Attenuator.Stream() * m.ModulationRange.Stream())
}
