package tools

import (
	"fmt"
	"io"

	"github.com/dalloriam/synthia/core"
)

type Wiretap struct {
	input core.Signal
	out   io.Writer
}

func Tap(inputSignal core.Signal, out io.Writer) *Wiretap {
	return &Wiretap{input: inputSignal, out: out}
}

func (w *Wiretap) Stream() float64 {
	if w.input == nil {
		return 0
	}

	val := w.input.Stream()

	if _, err := w.out.Write([]byte(fmt.Sprintf("%f\n", val))); err != nil {
		panic(err)
	}

	return val
}
