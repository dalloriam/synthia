package synthia

import "math"

type WaveShape int

const (
	SINE = iota
	SQUARE
)

type Oscillator struct {
	Frequency *Knob

	shape  WaveShape
	Volume byte

	phase   float64
	radians float64
}

func NewOscillator(freq float64, shape WaveShape) *Oscillator {
	osc := &Oscillator{
		shape:  shape,
		Volume: math.MaxUint8, // By default, oscillators output max volume and the mixer is tasked with vol. management
	}
	osc.Frequency = NewKnob(freq)
	return osc
}

func (o *Oscillator) Shape() WaveShape {
	return o.shape
}

func (o *Oscillator) ChangeShape(shape WaveShape) {
	o.shape = shape
}

func (o *Oscillator) incrementPhase() {
	if o.phase > 2*math.Pi {
		o.phase = o.phase - (2 * math.Pi)
	}
	o.phase += o.Frequency.Value() * math.Pi / float64(sampleRate)
}

func (o *Oscillator) sine(p []float64) {
	nbOfSamples := len(p)
	volFactor := float64(o.Volume) / float64(math.MaxUint8)

	for i := 0; i < nbOfSamples; i++ {
		o.incrementPhase()
		sin := math.Sin(o.phase)

		p[i] = sin * volFactor * math.MaxUint16 / 2
	}

}

func (o *Oscillator) square(p []float64) {
	nbOfSamples := len(p)

	var wv float64
	volFactor := float64(o.Volume) / float64(math.MaxUint8)

	for i := 0; i < nbOfSamples; i++ {
		o.incrementPhase()
		if math.Sin(o.phase) > 0 {
			wv = math.MaxUint16
		} else {
			wv = 0
		}
		p[i] = wv * volFactor
	}
}

func (o *Oscillator) Stream(p []float64) (int, error) {
	switch o.shape {
	case SINE:
		o.sine(p)
	case SQUARE:
		o.square(p)
	}

	return len(p), nil
}
