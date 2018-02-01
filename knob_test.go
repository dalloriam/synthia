package synthia_test

import (
	"testing"

	"github.com/dalloriam/synthia"
)

type mockLine struct {
	Val float64
}

func (m *mockLine) Stream(p []float64) {
	for i := 0; i < len(p); i++ {
		p[i] = m.Val
	}
}

func TestNewKnob(t *testing.T) {

	t.Run("initializes line to nil", func(t *testing.T) {
		k := synthia.NewKnob(0.0)

		if k.Line != nil {
			t.Errorf("doesn't initialize knob line to nil")
		}
	})

	t.Run("initializes value to proper default", func(t *testing.T) {
		val := 42.42

		k := synthia.NewKnob(val)

		buf := make([]float64, 1)

		k.Stream(buf)

		if buf[0] != val {
			t.Errorf("expected knob value to be %d, got %d", val, buf[0])
		}
	})
}

func TestKnob_SetValue(t *testing.T) {

	t.Run("properly sets value", func(t *testing.T) {
		k := synthia.NewKnob(0.0)

		newVal := 42.42

		k.SetValue(newVal)

		buf := make([]float64, 1)

		k.Stream(buf)

		if buf[0] != newVal {
			t.Errorf("expected knob value to be %d, got %d", newVal, buf[0])
		}
	})
}

func TestKnob_Stream(t *testing.T) {
	t.Run("returns value if no line connected", func(t *testing.T) {
		val := 18.0

		k := synthia.NewKnob(val)
		buf := make([]float64, 1)

		k.Stream(buf)

		if val != buf[0] {
			t.Errorf("expected value %d, got %d", val, buf[0])
		}
	})

	t.Run("streams from line if line connected", func(t *testing.T) {
		line := &mockLine{42}

		knob := synthia.NewKnob(18.0)

		buf := make([]float64, 1)

		knob.Line = line

		knob.Stream(buf)

		if buf[0] != line.Val {
			t.Errorf("expected value %d, got %d", line.Val, buf[0])
		}
	})

	t.Run("fills arbitrarily-sized buffer", func(t *testing.T) {

		for i := 0; i < 100; i++ {
			line := &mockLine{42}
			knob := synthia.NewKnob(18.0)
			knob.Line = line

			buf := make([]float64, i)
			knob.Stream(buf)
		}
	})
}
