package synthia

// A Speaker is the default output device for the synthesizer it's hardcoded for now, but will be customizable in the future
type Speaker struct {
	bufferSize int
	Input      AudioStream
	player     StreamOutput
	status     chan bool
}

// NewSpeaker returns an initialized speaker instance
func NewSpeaker(output StreamOutput, bufferSize int) *Speaker {

	return &Speaker{bufferSize: bufferSize, player: output}
}

func (s *Speaker) convert(floatOutput []float64, p []byte) {
	offset := 0

	for i := 0; i < len(floatOutput); i++ {
		v := uint16(floatOutput[i])

		buf := []byte{uint8(v & 0xff), uint8(v >> 8)}
		for j := 0; j < len(buf); j++ {
			p[offset+j] = buf[j]
		}
		offset += len(buf)
	}
}

func (s *Speaker) play() {
	stpChan := s.status
	for {
		select {
		default:
			buf := make([]float64, s.bufferSize/2)
			s.Input.Stream(buf)
			outBuf := make([]byte, s.bufferSize)
			s.convert(buf, outBuf)

			_, err := s.player.Write(outBuf)

			// TODO: Handler properly
			if err != nil {
				panic(err)
			}
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
