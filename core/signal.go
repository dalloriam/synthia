package core

// A Signal is a struct capable of producing a signal. Similar to a cable in an actual modular synthesizer.
type Signal interface {
	Stream() float64
}

// A SteroSignal represents a buffered stereo signal source. (Ex: an audio mixer)
type StereoSignal interface {
	Stream(l, r []float64)
}
