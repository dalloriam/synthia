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
	"github.com/dalloriam/synthia"
	"github.com/hajimehoshi/oto"
	"github.com/dalloriam/synthia/modular"
)

const (
	bufferSize = 8000 // Size of the audio buffer.
	sampleRate = 44100 // Audio sample rate.
	audioChannelCount = 2 // Stereo.
	mixerChannelCount = 3 // Three oscillators.
	bytesPerSample = 2 // We're generating 16-bit PCM audio, so two bytes per sample.
)

func main() {


	// Open connection to speaker. (Dependency to hajimehoshi/oto is optional but recommended 
	// for regular audio playback). Since Synthia is a dependency-free library, the 
	// synthesizer will happily output sound to any struct that satisfies the io.Writer 
	// interface.
	soundOutput, err := oto.NewPlayer(sampleRate, audioChannelCount, bytesPerSample, bufferSize)
	if err != nil {
		panic(err)
	}

    // Create the synthesizer with three mixer channels and set it to output to our speakers.
	synth := synthia.NewSynth(mixerChannelCount, bufferSize, soundOutput)
	
	// Create an oscillator and set it to 220Hz.
	osc1 := modular.NewOscillator()
	osc1.Frequency.SetValue(220)

	// Map three different waves to the three outputs of our mixer.
	synth.Mixer.Channels[0].Input = osc1.Triangle
	synth.Mixer.Channels[1].Input = osc1.Sine
	synth.Mixer.Channels[2].Input = osc1.Saw

	// Block until terminated
	select{}
}

```