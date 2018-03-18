package modular

import (
	"math"

	"github.com/dalloriam/synthia/core"
)

type stage int

const (
	STAGE_OFF stage = iota
	STAGE_ATTACK
	STAGE_DECAY
	STAGE_SUSTAIN
	STAGE_RELEASE
)

type Envelope struct {
	currentStage stage
	lastOutValue float64

	// Trigger info
	lastTrigger    float64
	currentTrigger float64

	// Knobs
	CurveRatio, Attack, Decay, Sustain, Release *core.Knob

	// curve info
	base, multiplier float64

	Trigger core.Signal
}

func NewEnvelope() *Envelope {
	return &Envelope{
		currentStage:   STAGE_OFF,
		lastOutValue:   0,
		lastTrigger:    0,
		currentTrigger: 0,

		CurveRatio: core.NewKnob(0.01),
		Attack:     core.NewKnob(50),
		Decay:      core.NewKnob(50),
		Sustain:    core.NewKnob(0.5),
		Release:    core.NewKnob(50),

		base:       0,
		multiplier: 0,
	}
}

func (e *Envelope) off() float64 {
	if e.lastTrigger <= 0 && e.currentTrigger > 0 {
		// Switch to attack state
		e.base, e.multiplier = computeSlope(
			e.CurveRatio.Stream(),
			e.Attack.Stream(),
			1,
			false,
		)
		e.currentStage = STAGE_ATTACK
	}
	return 0
}

func (e *Envelope) attack() float64 {
	val := e.base + e.multiplier*e.lastOutValue

	if val >= 1 {
		// Switch to decay state
		ratio := e.CurveRatio.Stream()
		decay := e.Decay.Stream()
		sustain := e.Sustain.Stream()
		e.base, e.multiplier = computeSlope(
			ratio,
			decay,
			sustain,
			true,
		)
		e.currentStage = STAGE_DECAY
	}

	return val
}

func (e *Envelope) decay() float64 {
	val := e.base + e.multiplier*e.lastOutValue

	if val <= e.Sustain.Stream() {
		// Switch to release or hold (depending if still triggered)
		if e.currentTrigger > 0 {
			// Hold
			e.currentStage = STAGE_SUSTAIN
		} else {
			// Release
			e.base, e.multiplier = computeSlope(
				e.CurveRatio.Stream(),
				e.Release.Stream(),
				0,
				true,
			)
			e.currentStage = STAGE_RELEASE
		}
	}

	return val
}

func (e *Envelope) sustain() float64 {
	if e.currentTrigger <= 0 {
		// Switch to release state
		e.base, e.multiplier = computeSlope(
			e.CurveRatio.Stream(),
			e.Release.Stream(),
			0,
			true,
		)
		e.currentStage = STAGE_RELEASE
	}
	return e.lastOutValue
}

func (e *Envelope) release() float64 {
	if e.lastTrigger <= 0 && e.currentTrigger > 0 {
		// Switch to attack state
		e.base, e.multiplier = computeSlope(
			e.CurveRatio.Stream(),
			e.Attack.Stream(),
			1,
			false,
		)
		e.currentStage = STAGE_ATTACK
	}

	val := e.base + e.lastOutValue*e.multiplier

	if val < math.SmallestNonzeroFloat64 {
		e.currentStage = STAGE_OFF
	}

	return val
}

func (e *Envelope) Stream() float64 {
	var out float64

	e.lastTrigger = e.currentTrigger
	e.currentTrigger = e.Trigger.Stream()

	switch e.currentStage {
	case STAGE_OFF:
		out = e.off()
	case STAGE_ATTACK:
		out = e.attack()
	case STAGE_DECAY:
		out = e.decay()
	case STAGE_SUSTAIN:
		out = e.sustain()
	case STAGE_RELEASE:
		out = e.release()
	default:
		out = 0
	}

	e.lastOutValue = out
	return out
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
