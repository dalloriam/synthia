package modular

import (
	"reflect"
	"testing"

	"github.com/dalloriam/synthia/core"
)

func TestNewEnvelope(t *testing.T) {
	tests := []struct {
		name string
		want *Envelope
	}{
		{
			name: "returns correct envelope",
			want: &Envelope{
				currentStage:   StageOff,
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnvelope(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnvelope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_off(t *testing.T) {
	type fields struct {
		currentStage   stage
		lastOutValue   float64
		lastTrigger    float64
		currentTrigger float64
		CurveRatio     *core.Knob
		Attack         *core.Knob
		Decay          *core.Knob
		Sustain        *core.Knob
		Release        *core.Knob
		base           float64
		multiplier     float64
		Trigger        core.Signal
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "does nothing if not triggered",
			fields: fields{
				lastTrigger:    0,
				currentTrigger: 0,
				CurveRatio:     core.NewKnob(0.5),
				Attack:         core.NewKnob(0.5),
			},
			want: 0.0,
		},
		{
			name: "returns 0 when triggered",
			fields: fields{
				lastTrigger:    0,
				currentTrigger: 10,
				CurveRatio:     core.NewKnob(0.5),
				Attack:         core.NewKnob(0.5),
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				currentStage:   tt.fields.currentStage,
				lastOutValue:   tt.fields.lastOutValue,
				lastTrigger:    tt.fields.lastTrigger,
				currentTrigger: tt.fields.currentTrigger,
				CurveRatio:     tt.fields.CurveRatio,
				Attack:         tt.fields.Attack,
				Decay:          tt.fields.Decay,
				Sustain:        tt.fields.Sustain,
				Release:        tt.fields.Release,
				base:           tt.fields.base,
				multiplier:     tt.fields.multiplier,
				Trigger:        tt.fields.Trigger,
			}
			if got := e.off(); got != tt.want {
				t.Errorf("Envelope.off() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_attack(t *testing.T) {
	type fields struct {
		currentStage   stage
		lastOutValue   float64
		lastTrigger    float64
		currentTrigger float64
		CurveRatio     *core.Knob
		Attack         *core.Knob
		Decay          *core.Knob
		Sustain        *core.Knob
		Release        *core.Knob
		base           float64
		multiplier     float64
		Trigger        core.Signal
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				currentStage:   tt.fields.currentStage,
				lastOutValue:   tt.fields.lastOutValue,
				lastTrigger:    tt.fields.lastTrigger,
				currentTrigger: tt.fields.currentTrigger,
				CurveRatio:     tt.fields.CurveRatio,
				Attack:         tt.fields.Attack,
				Decay:          tt.fields.Decay,
				Sustain:        tt.fields.Sustain,
				Release:        tt.fields.Release,
				base:           tt.fields.base,
				multiplier:     tt.fields.multiplier,
				Trigger:        tt.fields.Trigger,
			}
			if got := e.attack(); got != tt.want {
				t.Errorf("Envelope.attack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_decay(t *testing.T) {
	type fields struct {
		currentStage   stage
		lastOutValue   float64
		lastTrigger    float64
		currentTrigger float64
		CurveRatio     *core.Knob
		Attack         *core.Knob
		Decay          *core.Knob
		Sustain        *core.Knob
		Release        *core.Knob
		base           float64
		multiplier     float64
		Trigger        core.Signal
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				currentStage:   tt.fields.currentStage,
				lastOutValue:   tt.fields.lastOutValue,
				lastTrigger:    tt.fields.lastTrigger,
				currentTrigger: tt.fields.currentTrigger,
				CurveRatio:     tt.fields.CurveRatio,
				Attack:         tt.fields.Attack,
				Decay:          tt.fields.Decay,
				Sustain:        tt.fields.Sustain,
				Release:        tt.fields.Release,
				base:           tt.fields.base,
				multiplier:     tt.fields.multiplier,
				Trigger:        tt.fields.Trigger,
			}
			if got := e.decay(); got != tt.want {
				t.Errorf("Envelope.decay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_sustain(t *testing.T) {
	type fields struct {
		currentStage   stage
		lastOutValue   float64
		lastTrigger    float64
		currentTrigger float64
		CurveRatio     *core.Knob
		Attack         *core.Knob
		Decay          *core.Knob
		Sustain        *core.Knob
		Release        *core.Knob
		base           float64
		multiplier     float64
		Trigger        core.Signal
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				currentStage:   tt.fields.currentStage,
				lastOutValue:   tt.fields.lastOutValue,
				lastTrigger:    tt.fields.lastTrigger,
				currentTrigger: tt.fields.currentTrigger,
				CurveRatio:     tt.fields.CurveRatio,
				Attack:         tt.fields.Attack,
				Decay:          tt.fields.Decay,
				Sustain:        tt.fields.Sustain,
				Release:        tt.fields.Release,
				base:           tt.fields.base,
				multiplier:     tt.fields.multiplier,
				Trigger:        tt.fields.Trigger,
			}
			if got := e.sustain(); got != tt.want {
				t.Errorf("Envelope.sustain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_release(t *testing.T) {
	type fields struct {
		currentStage   stage
		lastOutValue   float64
		lastTrigger    float64
		currentTrigger float64
		CurveRatio     *core.Knob
		Attack         *core.Knob
		Decay          *core.Knob
		Sustain        *core.Knob
		Release        *core.Knob
		base           float64
		multiplier     float64
		Trigger        core.Signal
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				currentStage:   tt.fields.currentStage,
				lastOutValue:   tt.fields.lastOutValue,
				lastTrigger:    tt.fields.lastTrigger,
				currentTrigger: tt.fields.currentTrigger,
				CurveRatio:     tt.fields.CurveRatio,
				Attack:         tt.fields.Attack,
				Decay:          tt.fields.Decay,
				Sustain:        tt.fields.Sustain,
				Release:        tt.fields.Release,
				base:           tt.fields.base,
				multiplier:     tt.fields.multiplier,
				Trigger:        tt.fields.Trigger,
			}
			if got := e.release(); got != tt.want {
				t.Errorf("Envelope.release() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Stream(t *testing.T) {
	type fields struct {
		currentStage   stage
		lastOutValue   float64
		lastTrigger    float64
		currentTrigger float64
		CurveRatio     *core.Knob
		Attack         *core.Knob
		Decay          *core.Knob
		Sustain        *core.Knob
		Release        *core.Knob
		base           float64
		multiplier     float64
		Trigger        core.Signal
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				currentStage:   tt.fields.currentStage,
				lastOutValue:   tt.fields.lastOutValue,
				lastTrigger:    tt.fields.lastTrigger,
				currentTrigger: tt.fields.currentTrigger,
				CurveRatio:     tt.fields.CurveRatio,
				Attack:         tt.fields.Attack,
				Decay:          tt.fields.Decay,
				Sustain:        tt.fields.Sustain,
				Release:        tt.fields.Release,
				base:           tt.fields.base,
				multiplier:     tt.fields.multiplier,
				Trigger:        tt.fields.Trigger,
			}
			if got := e.Stream(); got != tt.want {
				t.Errorf("Envelope.Stream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeSlope(t *testing.T) {
	type args struct {
		ratio  float64
		length float64
		tgt    float64
		isExp  bool
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := computeSlope(tt.args.ratio, tt.args.length, tt.args.tgt, tt.args.isExp)
			if got != tt.want {
				t.Errorf("computeSlope() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("computeSlope() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
