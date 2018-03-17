package synthia

import "io"

// A Speaker is the output device for the synthesizer.
type Speaker struct {
	bufferSize int
	Input      StereoSignal
	player     io.Writer
	status     chan bool
}

// NewSpeaker returns an initialized speaker instance
func NewSpeaker(output io.Writer, bufferSize int) *Speaker {

	return &Speaker{bufferSize: bufferSize, player: output}
}

func (s *Speaker) convert(rightIn, leftIn []float64, p []byte) {
	offset := 0

	inLength := len(rightIn)

	chans := [][]float64{leftIn, rightIn}

	for i := 0; i < inLength; i++ {

		for chanIdx := 0; chanIdx < len(chans); chanIdx++ {
			v := uint16(chans[chanIdx][i])

			buf := []byte{uint8(v & 0xff), uint8(v >> 8)}
			for j := 0; j < len(buf); j++ {
				p[offset+j] = buf[j]
			}

			offset += len(buf)

		}
	}
}

func (s *Speaker) processBuffer() {
	rightBuf := make([]float64, s.bufferSize)
	leftBuf := make([]float64, s.bufferSize)

	for i := 0; i < s.bufferSize; i++ {

	}

	s.Input.Stream(leftBuf, rightBuf)

	outBuf := make([]byte, s.bufferSize*4)
	s.convert(rightBuf, leftBuf, outBuf)

	_, err := s.player.Write(outBuf)

	// TODO: Handler properly
	if err != nil {
		panic(err)
	}
}

func (s *Speaker) play() {
	stpChan := s.status
	for {
		select {
		default:
			s.processBuffer()
		case <-stpChan:
			return
		}
	}
}

// Start starts the speaker
func (s *Speaker) Start() {
	s.status = make(chan bool)
	go s.play()
}

// Stop stops the speaker
func (s *Speaker) Stop() {
	s.status <- true
	close(s.status)
}
