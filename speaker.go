package synthia

// A Speaker is the output device for the synthesizer.
type Speaker struct {
	bufferSize int
	chunks     int
	Input      StereoSignal
	status     chan bool
}

// NewSpeaker returns an initialized speaker instance
func NewSpeaker(bufferSize, framesize int) *Speaker {
	return &Speaker{bufferSize: bufferSize, chunks: int(float64(framesize) / float64(bufferSize))}
}

func (s *Speaker) ProcessBuffer(in []float32, out [][]float32) {
	rightBuf := make([]float64, s.bufferSize)
	leftBuf := make([]float64, s.bufferSize)

	s.Input.Stream(leftBuf, rightBuf)

	for k := 0; k < s.chunks; k++ {
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
