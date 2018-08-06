package core_test

import (
	"testing"

	"github.com/dalloriam/synthia/core"
)

type mockLine struct {
	Val float64
}

func (m *mockLine) Stream() float64 {
	return m.Val
}

func TestNewKnob(t *testing.T) {

	t.Run("initializes line to nil", func(t *testing.T) {
		k := core.NewKnob(0.0)

		if k.Line != nil {
			t.Errorf("doesn't initialize knob line to nil")
		}
	})

	t.Run("initializes value to proper default", func(t *testing.T) {
		expected := 42.42

		k := core.NewKnob(expected)

		actual := k.Stream()

		if actual != expected {
			t.Errorf("expected knob value to be %f, got %f", expected, actual)
		}
	})
}

func TestKnob_SetValue(t *testing.T) {

	t.Run("properly sets value", func(t *testing.T) {
		k := core.NewKnob(0.0)

		newVal := 42.42

		k.SetValue(newVal)

		actual := k.Stream()

		if actual != newVal {
			t.Errorf("expected knob value to be %f, got %f", newVal, actual)
		}
	})
}

func TestKnob_Stream(t *testing.T) {
	t.Run("returns value if no line connected", func(t *testing.T) {
		val := 18.0

		k := core.NewKnob(val)

		actual := k.Stream()

		if actual != val {
			t.Errorf("expected value %f, got %f", val, actual)
		}
	})

	t.Run("streams from line if line connected", func(t *testing.T) {
		line := &mockLine{42}

		knob := core.NewKnob(18.0)

		knob.Line = line

		actual := knob.Stream()

		if actual != line.Val {
			t.Errorf("expected value %f, got %f", line.Val, actual)
		}
	})
}

func TestKnob_GetValue(t *testing.T) {
	t.Run("returns last set value when no line connected", func(t *testing.T) {
		oldValue := 14.0
		newValue := 36.0

		knob := core.NewKnob(oldValue)

		if knob.GetValue() != oldValue {
			t.Errorf("expected value %f, got %f", oldValue, knob.GetValue())
		}

		knob.SetValue(newValue)

		if knob.GetValue() != newValue {
			t.Errorf("expected value %f, got %f", newValue, knob.GetValue())
		}
	})

	t.Run("returns last streamed value when line connected", func(t *testing.T) {
		line := &mockLine{18}

		knob := core.NewKnob(42.0)

		knob.Line = line
		knob.Stream()

		if knob.GetValue() != line.Val {
			t.Errorf("expected value %f, got %f", line.Val, knob.GetValue())
		}
	})

	t.Run("doesnt mutate knob value or call stream", func(t *testing.T) {
		line := &mockLine{18}

		expected := 42.0

		knob := core.NewKnob(expected)

		knob.Line = line

		actual := knob.GetValue()

		if actual != expected {
			t.Errorf("expected value %f, got %f", expected, actual)
		}
	})
}
