package synthia

import "github.com/dalloriam/synthia/core"

// Synthia is the core synthesizer struct
type Synthia struct {
	bufferSize int
	chunkSize  int
	Mixer      *core.Mixer
	output     audioBackend
}

type audioBackend interface {
	Start(callback func(in []float32, out [][]float32)) error
	FrameSize() int
}

// New returns a new synthesizer with an already-initialized mixer
func New(channelCount, bufferSize int, output audioBackend) *Synthia {
	m := core.NewMixer(channelCount)

	synth := &Synthia{
		bufferSize: bufferSize,
		chunkSize:  int(float64(output.FrameSize()) / float64(bufferSize)),
		Mixer:      m,
		output:     output,
	}

	err := output.Start(synth.processBuffer)

	if err != nil {
		panic(err)
	}

	return synth
}

func (s *Synthia) processBuffer(in []float32, out [][]float32) {
	rightBuf := make([]float64, s.bufferSize)
	leftBuf := make([]float64, s.bufferSize)

	s.Mixer.Stream(leftBuf, rightBuf)

	for k := 0; k < s.chunkSize; k++ {
		offset := s.bufferSize * k

		for i := 0; i < len(out); i++ {
			for j := 0; j < s.bufferSize; j++ {
				if i%2 == 0 {
					out[i][j+offset] = float32(leftBuf[j])
				} else {
					out[i][j+offset] = float32(rightBuf[j])
				}
			}
		}

	}
}
