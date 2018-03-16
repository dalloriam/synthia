package synthia

// A Signal is a struct capable of producing a signal. Similar to a cable in an actual modular synthesizer.
type Signal interface {
	Stream(p []float64)
}

type StereoSignal interface {
	Stream(l, r []float64)
}
