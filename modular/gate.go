package modular

import "github.com/dalloriam/synthia/core"

// A Gate applies a volume multiplier from the Trigger to the Input signal.
type Gate struct {
	Input   core.Signal
	Trigger core.Signal
}

// NewGate returns an empty gate.
func NewGate() *Gate {
	return &Gate{}
}

// Stream streams a sample from the gate.
func (g *Gate) Stream() float64 {
	return g.Input.Stream() * g.Trigger.Stream()
}
