package modular

import (
	"math"

	"github.com/dalloriam/synthia/core/constants"

	"github.com/dalloriam/synthia/core"
)

type FilterType int

const (
	LowPassFilter FilterType = iota
	HighPassFilter
)

const q = 0.5

type Filter struct {
	Cutoff *core.Knob
	Input  core.Signal
	Type   FilterType

	ic1eq, ic2eq float64
}

func NewFilter(filterType FilterType) *Filter {
	filt := &Filter{
		Cutoff: core.NewKnob(500), // Default cutoff frequency is 500hz.
		Type:   filterType,
	}

	return filt
}

func (f *Filter) Stream() float64 {
	if f.Input == nil {
		return 0
	}
	g := tan((f.Cutoff.Stream() * 10 / constants.SampleRate) * math.Pi)
	k := 1 / q

	a1 := 1 / (1 + g*(g+k))
	a2 := g * a1
	a3 := g * a2

	var m0, m1, m2 float64

	switch f.Type {
	case LowPassFilter:
		m0 = 0
		m1 = 0
		m2 = 1
	case HighPassFilter:
		m0 = 1
		m1 = -k
		m2 = -1
	}

	v0 := f.Input.Stream()
	v3 := v0 - f.ic2eq
	v1 := a1 * f.ic1eq * a2 * v3
	v2 := f.ic2eq + a2*f.ic1eq + a3*v3

	f.ic1eq = 2*v1 - f.ic1eq
	f.ic2eq = 2*v2 - f.ic2eq

	return m0*v0 + m1*v1 + m2*v2
}
