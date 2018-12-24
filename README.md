# Synthia
[![Go Report Card](https://goreportcard.com/badge/github.com/dalloriam/synthia)](https://goreportcard.com/report/github.com/dalloriam/synthia)
[![GoDoc](https://godoc.org/github.com/dalloriam/synthia?status.svg)](https://godoc.org/github.com/dalloriam/synthia)
[![CircleCI](https://circleci.com/gh/dalloriam/synthia.svg?style=svg)](https://circleci.com/gh/dalloriam/synthia)

Synthia is a standalone (as in _dependency-free_) toolkit for synthesizing audio in Go. The project is still in its
early stages, but should improve at a good pace.

## Quick Start

```go
package main

import (
	"github.com/gordonklaus/portaudio"
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/synthia"
)

const (
	bufferSize = 512 // Size of the audio buffer.
	sampleRate = 44100 // Audio sample rate.
	audioChannelCount = 2 // Stereo.
	mixerChannelCount = 2 // Three oscillators.
)

type AudioBackend struct {
	params portaudio.StreamParameters
}

func (b *AudioBackend) Start(callback func(in []float32, out [][]float32)) error {
	stream, err := portaudio.OpenStream(b.params, callback)
	if err != nil {
		return err
	}

	return stream.Start()
}

func (b *AudioBackend) FrameSize() int {
	return b.params.FramesPerBuffer
}

func newBackend() *AudioBackend {
	// Quick-and-dirty way to initialize portaudio
	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}

	devices, err := portaudio.Devices()
	if err != nil {
		panic(err)
	}
	inDevice, outDevice := devices[0], devices[1]

	params := portaudio.LowLatencyParameters(inDevice, outDevice)

	params.Input.Channels = 1
	params.Output.Channels = audioChannelCount

	params.SampleRate = float64(sampleRate)
	params.FramesPerBuffer = bufferSize

	return &AudioBackend{params}
}

func main() {
	backend := newBackend()

	// Create a new oscillator.
	osc := modular.NewOscillator()

	// Instantiate a clock @ 120bpm.
	clock := modular.NewClock()
	clock.Tempo.SetValue(120)

	// Define a sequencer playing a C Major scale in quarter notes.
	seq := modular.NewSequencer([]float64{130.81, 146.83, 164.1, 174.61, 196, 220, 246.94, 261.63})
	seq.Clock = clock
	seq.BeatsPerStep.SetValue(0.50)

	// Create an ADSR envelope with a release time of 800ms.
	// Attack defaults to 0.5ms, decay to 50ms, and sustain to 0.5ms.
	attackEnvelope := modular.NewEnvelope()
	attackEnvelope.Release.SetValue(800)

	// Modulate the frequency of the oscillator with the defined sequencer.
	osc.Frequency.Line = seq

	// Modulate the volume of the oscillator with the envelope.
	osc.Volume.Line = attackEnvelope

	// Trigger the envelope on every sequencer note change.
	attackEnvelope.Trigger = seq.Trigger

	// Define a high-pass filter taking in the Sawtooth output of the oscillator.
	filter := modular.NewFilter(modular.HighPassFilter)
	filter.Input = osc.Saw

	// Define an LFO running from 300 to 2000 @ 1Hz.
	filterLfo := modular.NewLFO(2000, 300)
	filterLfo.Osc.Frequency.SetValue(1)

	// Modulate the filter cutoff with the LFO.
	filter.Cutoff.Line = filterLfo
	
	// Create a synthesizer instance.
	synth := synthia.New(mixerChannelCount, bufferSize, backend)

	// Patch the filter output to mixer channel 1.
	synth.Mixer.Channels[0].Input = filter

	// Wait until user exit.
	select {}
}
```