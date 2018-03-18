package modular

/* MIT License

Copyright (c) 2017 Brett Buddin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Original source at: https://github.com/brettbuddin/shaden/blob/master/dsp/trig.go
*/

import "math"

const (
	sineLength = 1024
	sineStep   = sineLength / (2 * math.Pi)
)

var (
	sineTable = make([]float64, sineLength)
	sineDiff  = make([]float64, sineLength)
)

func init() {
	for i := 0; i < sineLength; i++ {
		sineTable[i] = math.Sin(float64(i) * (1 / sineStep))
	}
	for i := 0; i < sineLength; i++ {
		next := sineTable[(i+1)%sineLength]
		sineDiff[i] = next - sineTable[i]
	}
}

func sin(x float64) float64 {
	step := x * sineStep
	if x < 0 {
		step = -step
	}

	var (
		trunc = int(step)
		i     = trunc % sineLength
		out   = sineTable[i] + sineDiff[i]*(step-float64(trunc))
	)

	if x < 0 {
		return -out
	}
	return out
}

// tan is a lookup table version of math.Tan
func tan(x float64) float64 {
	return sin(x) / cos(x)
}

// cos is a lookup table version of math.Cos
func cos(x float64) float64 {
	return sin(x + 0.5*math.Pi)
}
