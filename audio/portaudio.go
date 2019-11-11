package audio

import (
	"github.com/dalloriam/synthia/core/constants"
	"github.com/gordonklaus/portaudio"
)

type PortaudioBackend struct {
	params portaudio.StreamParameters
}

func NewPortaudioBackend(bufferSize, inputDevice, outputDevice int) (*PortaudioBackend, error) {
	// Quick-and-dirty way to initialize portaudio
	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}

	devices, err := portaudio.Devices()
	if err != nil {
		return nil, err
	}
	// TODO: Detect devices properly
	inDevice, outDevice := devices[inputDevice], devices[outputDevice]

	params := portaudio.LowLatencyParameters(inDevice, outDevice)

	params.Input.Channels = 1
	params.Output.Channels = constants.ChannelCount

	params.SampleRate = float64(constants.SampleRate)
	params.FramesPerBuffer = bufferSize

	return &PortaudioBackend{params}, err
}

func (b *PortaudioBackend) Start(callback func(in []float32, out [][]float32)) error {
	stream, err := portaudio.OpenStream(b.params, callback)
	if err != nil {
		return err
	}

	return stream.Start()
}

func (b *PortaudioBackend) FrameSize() int {
	return b.params.FramesPerBuffer
}
