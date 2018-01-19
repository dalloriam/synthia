package synthia

type AudioStream interface {
	Stream(p []float64)
}
