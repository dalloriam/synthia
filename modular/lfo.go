package modular

// An LFO is a Low-Frequency Oscillator. It differs from the actual oscillator in that it allows for setting a floor and a ceiling.
type LFO struct {
	Oscillator
	MaxValue float64
	MinValue float64
	Shape    WaveShape
}

// NewLFO returns an new LFO.
func NewLFO() *LFO {
	l := &LFO{
		Oscillator: *NewOscillator(),
		MaxValue:   1,
		MinValue:   -1,
		Shape:      SINE,
	}
	l.Frequency.SetValue(4)
	return l
}

// Stream returns the current LFO phase.
func (l *LFO) Stream() float64 {
	line := l.GetOutput(l.Shape)
	sample := line.Stream() + 1
	return ((sample / 2) * (l.MaxValue - l.MinValue)) + l.MinValue
}
