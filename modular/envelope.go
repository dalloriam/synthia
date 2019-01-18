// Much of the magic here provided by
// https://github.com/brettbuddin/shaden/blob/master/unit/adsr.go
package modular

import (
	"math"

	"github.com/dalloriam/synthia/core/constants"

	"github.com/dalloriam/synthia/core"
)

type stage int

// These stages represent the different envelope stages.
const (
	StageOff stage = iota
	StageAttack
	StageDecay
	StageSustain
	StageRelease
)

// An Envelope represents an ADSR envelope.
type Envelope struct {
	currentStage stage
	lastOutValue float64

	lastSustain float64 // Buffer for sustain value to preserve 1:1 sample ratio

	// Trigger info
	lastTrigger    float64
	currentTrigger float64

	// Knobs
	CurveRatio, Attack, Decay, Sustain, Release *core.Knob

	// curve info
	base, multiplier float64

	Trigger core.Signal
}

// NewEnvelope returns a new ADSR envelope generator.
func NewEnvelope() *Envelope {
	return &Envelope{
		currentStage:   StageOff,
		lastOutValue:   0,
		lastTrigger:    0,
		currentTrigger: 0,

		CurveRatio: core.NewKnob(0.01),
		Attack:     core.NewKnob(0.5),
		Decay:      core.NewKnob(50),
		Sustain:    core.NewKnob(0.5),
		Release:    core.NewKnob(50),

		lastSustain: 0,

		base:       0,
		multiplier: 0,
	}
}

func (e *Envelope) msToSamples(msCount float64) float64 {
	return msCount * constants.SampleRate * 0.001
}

func (e *Envelope) updateBaseMult() {
	switch e.currentStage {
	case StageAttack:
		e.base, e.multiplier = computeSlope(
			e.CurveRatio.Stream(),
			e.msToSamples(e.Attack.Stream()),
			1,
			false,
		)
	case StageDecay:
		ratio := e.CurveRatio.Stream()
		decay := e.msToSamples(e.Decay.Stream())
		sustain := e.msToSamples(e.Sustain.Stream())
		e.lastSustain = sustain
		e.base, e.multiplier = computeSlope(
			ratio,
			decay,
			sustain,
			true,
		)
	case StageRelease:
		e.base, e.multiplier = computeSlope(
			e.CurveRatio.Stream(),
			e.msToSamples(e.Release.Stream()),
			0,
			true,
		)
	}
}

func (e *Envelope) off() float64 {
	if e.lastTrigger <= 0 && e.currentTrigger > 0 {
		// Switch to attack state
		e.currentStage = StageAttack
	}
	return 0
}

func (e *Envelope) attack() float64 {
	e.updateBaseMult()
	val := e.base + e.multiplier*e.lastOutValue

	if val >= 1 {
		// Switch to decay state
		e.currentStage = StageDecay
	}

	return val
}

func (e *Envelope) decay() float64 {
	e.updateBaseMult()
	val := e.base + e.multiplier*e.lastOutValue

	if val <= e.msToSamples(e.lastSustain) {
		// Switch to release or hold (depending if still triggered)
		if e.currentTrigger > 0 {
			// Hold
			e.currentStage = StageSustain
		} else {
			// Release
			e.currentStage = StageRelease
		}
	}

	return val
}

func (e *Envelope) sustain() float64 {
	e.updateBaseMult()
	if e.currentTrigger <= 0 {
		e.currentStage = StageRelease
	}
	return e.lastOutValue
}

func (e *Envelope) release() float64 {
	e.updateBaseMult()
	if e.lastTrigger <= 0 && e.currentTrigger > 0 {
		e.currentStage = StageAttack
	}

	val := e.base + e.lastOutValue*e.multiplier

	if val < math.SmallestNonzeroFloat64 {
		e.currentStage = StageOff
		return 0
	}

	return val
}

// Stream returns the current envelope phase.
func (e *Envelope) Stream() float64 {
	var out float64

	e.lastTrigger = e.currentTrigger
	e.currentTrigger = e.Trigger.Stream()

	switch e.currentStage {
	case StageOff:
		out = e.off()
	case StageAttack:
		out = e.attack()
	case StageDecay:
		out = e.decay()
	case StageSustain:
		out = e.sustain()
	case StageRelease:
		out = e.release()
	default:
		out = 0
	}

	e.lastOutValue = out
	return math.Min(out, 1.0)
}

func computeSlope(ratio, length, tgt float64, isExp bool) (float64, float64) {

	var base, mult float64
	mult = math.Exp(-math.Log((1+ratio)/ratio) / length)

	if isExp {
		ratio = -ratio
	}

	base = (tgt + ratio) * (1.0 - mult)
	return base, mult
}
