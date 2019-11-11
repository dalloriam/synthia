package main

import (
	"github.com/dalloriam/synthia/audio"
	"github.com/dalloriam/synthia/core"
	"github.com/dalloriam/synthia/tools"
)

const (
	bufferSize   = 512
	channelCount = 2
)

func main() {
	// Initialize the portaudio backend.
	backend, err := audio.NewPortaudioBackend(bufferSize, 0, 0)
	if err != nil {
		panic(err)
	}

	synth := core.NewSynth(channelCount, bufferSize, backend)
	patchDef, err := tools.LoadPatchFile("patch.json")
	if err != nil {
		panic(err)
	}
	if err := tools.Patch(patchDef, synth); err != nil {
		panic(err)
	}

	select {}
}
