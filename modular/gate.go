package modular

type Gate struct {
	Opened bool
}

// NewGate returns an empty gate.
func NewGate() *Gate {
	return &Gate{Opened: false}
}

// Stream streams a sample from the gate.
func (g *Gate) Stream() float64 {
	if g.Opened {
		return 1
	} else {
		return 0
	}
}
