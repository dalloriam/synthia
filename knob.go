package synthia

// A Knob represents a module parameter that can be manually tweaked as well as plugged in to a line
type Knob struct {
	Line  Signal
	value float64
}

// NewKnob returns a knob
func NewKnob(defaultVal float64) *Knob {
	k := &Knob{
		Line:  nil,
		value: defaultVal,
	}
	return k
}

// SetValue alters the internal value of the knob (equivalent to actually turning a knob on a real synthesizer module)
func (k *Knob) SetValue(val float64) {
	k.value = val
}

// Stream fills the buffer with the current knob value if no line connected. If a line is connected to the knob, it reads from the line instead
func (k *Knob) Stream() float64 {
	if k.Line == nil {
		return k.value
	}
	return k.Line.Stream()
}
