package modular

import (
	"github.com/dalloriam/synthia/core"
	"github.com/pkg/errors"
)

type WaveTable struct {
	Data [][]float64

	Frequency *core.Knob
	Position  *core.Knob
	Volume    *core.Knob

	generator *toneGenerator
}

func NewWaveTable(data [][]float64) (*WaveTable, error) {
	if len(data) == 0 {
		return nil, errors.New("wave table cannot be empty")
	}

	wt := &WaveTable{
		Data:      data,
		Frequency: core.NewKnob(440),
		Volume:    core.NewKnob(1),
		Position:  core.NewKnob(0),
	}

	wt.generator = &toneGenerator{phase: 0, volume: wt.Volume, frequency: wt.Frequency, tone: wt.tone}

	return wt, nil
}

func (wt *WaveTable) tone(phase float64) float64 {
	wtPos := int(wt.Position.Stream())
	return wt.Data[wtPos][int(phase/twoPi*float64(len(wt.Data[wtPos])))]
}

// Stream returns the current value at the current wavetable position.
func (wt *WaveTable) Stream() float64 {
	return wt.generator.Stream()
}
