package core

import (
	"sync"
)

type allocationPool struct {
	bufPool sync.Pool
}

func newAllocationPool(bufferSize int) *allocationPool {
	return &allocationPool{bufPool: sync.Pool{New: func() interface{} { return make([]float64, bufferSize) }}}
}

func (a *allocationPool) Get() []float64 {
	return a.bufPool.Get().([]float64)
}

func (a *allocationPool) Put(x []float64) {
	for i := 0; i < len(x); i++ {
		x[i] = 0
	}
	a.bufPool.Put(x)
}
