package main

import (
	"sync"
	"testing"
)

var reqSlicePool = sync.Pool{
	New: func() interface{} {
		var sl = make([]byte, 0)
		return &sl
	},
}

type reqStruct struct {
	userData []byte
}

func handleRequestWithPool(sliceLen int) {
	var userData = reqSlicePool.Get().(*[]byte)
	defer reqSlicePool.Put(userData)

	if len(*userData) < sliceLen {
		*userData = make([]byte, sliceLen)
	}

	// 往 userData 里塞一些数据，这里就省略了

	var req = reqStruct{userData: *userData}
	_ = req
}

func handleRequestWithoutPool(sliceLen int) {
	var userData = make([]byte, sliceLen)
	// 往 userData 里塞一些数据，这里就省略了

	var req = reqStruct{userData: userData}
	_ = req
}

func BenchmarkZeroGarbage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		handleRequestWithPool(10)
	}
}

func BenchmarkHasGarbage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		handleRequestWithoutPool(10)
	}
}
