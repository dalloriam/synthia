package modular

import (
	"math"

	"github.com/dalloriam/synthia/core"
)

type coefficients struct {
	A1, A2, A3, M0, M1, M2 float64
	A, ASqrt               float64
}

func newCoefficients() *coefficients {
	return &coefficients{A: 1, ASqrt: 1}
}

func (c *coefficients) computeA(g, k float64) {
	c.A1 = 1 / (1 + g*(g+k))
	c.A2 = g * c.A1
	c.A3 = g * c.A2
}

func (c *coefficients) computeK(q float64, useGain bool) float64 {
	var denom float64
	if useGain {
		denom = q * c.A
	} else {
		denom = q
	}

	return 1 / denom
}

func (c *coefficients) Update(cutoffFrequency, q float64, filterType string) {
	g := tan((cutoffFrequency / sampleRate) * math.Pi)
	k := c.computeK(q, false) // TODO: Support gain

	switch filterType {
	case "lowpass":
		c.computeA(g, k)
		c.M0 = 0
		c.M1 = 0
		c.M2 = 1
	}
}

type Filter struct {
	Cutoff *core.Knob
	Input  core.Signal

	coeff        *coefficients
	v1, v2, v3   float64
	ic1eq, ic2eq float64
}

func NewFilter() *Filter {
	filt := &Filter{
		Cutoff: core.NewKnob(500), // Default cutoff frequency is 500hz.
		coeff:  newCoefficients(),
	}

	return filt
}

func (f *Filter) lpf(x, alpha, x0 float64) float64 {
	return alpha * (x - x0)
}

func (f *Filter) Stream() float64 {
	if f.Input == nil {
		return 0
	}
	f.coeff.Update(f.Cutoff.Stream()*10, 1, "lowpass")

	v0 := f.Input.Stream()
	f.v3 = v0 - f.ic2eq
	f.v1 = f.coeff.A1 * f.ic1eq * f.coeff.A2 * f.v3
	f.v2 = f.ic2eq + f.coeff.A2*f.ic1eq + f.coeff.A3*f.v3

	f.ic1eq = 2*f.v1 - f.ic1eq
	f.ic2eq = 2*f.v2 - f.ic2eq

	return f.coeff.M0*v0 + f.coeff.M1*f.v1 + f.coeff.M2*f.v2
}
