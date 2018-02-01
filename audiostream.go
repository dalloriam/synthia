package synthia

// An AudioStream is a struct capable of producing a signal. Similar to a cable in an actual modular synthesizer.
type AudioStream interface {
	Stream(p []float64)
}
