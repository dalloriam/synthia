package synthia

type AudioStream interface {
	Stream(p []float64) (int, error)
}
