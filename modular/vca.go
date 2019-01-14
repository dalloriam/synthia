package modular

import (
	"github.com/dalloriam/synthia/core"
)

type VCA struct {
	Input  core.Signal
	Gate   core.Signal
	Volume *core.Knob
}

func NewVCA() *VCA {
	return &VCA{Volume: core.NewKnob(1)}
}

func (v *VCA) Stream() float64 {
	if v.Input == nil {
		return 0
	}
	inputVal := v.Input.Stream()
	volVal := v.Volume.Stream()

	if v.Gate == nil || v.Gate.Stream() > 0 {
		return inputVal * volVal
	}

	return 0
}
