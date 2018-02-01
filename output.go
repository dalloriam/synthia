package synthia

type StreamOutput interface {
	Write(data []uint8) (int, error)
}
