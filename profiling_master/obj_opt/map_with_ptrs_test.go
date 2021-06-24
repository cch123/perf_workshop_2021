package main

import "testing"

func BenchmarkMapWithoutPtrs(b *testing.B) {
	for i:=0;i< b.N;i++{
		var m = make(map[int]int)
		for i:=0;i<10;i++ {
			m[i] = i
		}
	}
}

func BenchmarkMapWithPtrs(b *testing.B) {
	for i:=0;i< b.N;i++{
		var m = make(map[int]*int)
		for i:=0;i<10;i++ {
			var v = i
			m[i] = &v
		}
	}
}

