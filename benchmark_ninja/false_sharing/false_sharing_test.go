package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type nopad struct {
	x uint64
	y uint64
}

type withpad struct {
	x    uint64
	_pad [7]uint64
	y    uint64
}

func BenchmarkNoPad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var np nopad
		var wg sync.WaitGroup
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				for {
					if atomic.AddUint64(&np.x, 1) >= 10000000 {
						return
					}
				}
			}()
		}
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				for {
					if atomic.AddUint64(&np.y, 1) >= 10000000 {
						return
					}
				}
			}()
		}
		wg.Wait()
	}
}

func BenchmarkWithPad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var np withpad
		var wg sync.WaitGroup
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				for {
					if atomic.AddUint64(&np.x, 1) >= 10000000 {
						return
					}
				}
			}()
		}
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				for {
					if atomic.AddUint64(&np.y, 1) >= 10000000 {
						return
					}
				}
			}()
		}
		wg.Wait()
	}
}
