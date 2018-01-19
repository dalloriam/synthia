package synthia

type Knob struct {
	callbacks []func()
	Line      AudioStream
	value     float64
}

func NewKnob(defaultVal float64, callbacks ...func()) *Knob {
	k := &Knob{
		callbacks: callbacks,
		Line:      nil,
		value:     defaultVal,
	}
	return k
}

func (k *Knob) SetValue(val float64) {
	k.value = val

	for _, cb := range k.callbacks {
		cb()
	}
}

func (k *Knob) Value() float64 {
	if k.Line == nil {
		return k.value
	}
	buf := make([]float64, 1)
	k.Line.Stream(buf)
	return buf[0]
}
