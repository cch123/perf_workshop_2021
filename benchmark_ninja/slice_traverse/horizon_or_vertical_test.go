package main

import "testing"

var arr = make([][]int, 10000)

func init() {
	for i := 0; i < 10000; i++ {
		arr[i] = make([]int, 10000)
	}
}

func BenchmarkHorizontal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < len(arr); x++ {
			for y := 0; y < len(arr); y++ {
				arr[x][y] = 1
			}
		}
	}
}

func BenchmarkVertical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < len(arr); x++ {
			for y := 0; y < len(arr); y++ {
				arr[y][x] = 1
			}
		}
	}
}
