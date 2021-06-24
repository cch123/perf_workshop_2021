package main

import (
	"fmt"
	"testing"
)

func BenchmarkPlusConcat(b *testing.B) {
	var main string
	for i := 0; i < b.N; i++ {
		main = ""
		main += "userid : " + "1"
		main += "localtion : " + "ab"
	}
}

func BenchmarkSprintf(b *testing.B) {
	var main string
	for i := 0; i < b.N; i++ {
		main = ""
		main += fmt.Sprintf("userid : %v", "1")
		main += fmt.Sprintf("location : %v", "ab")
	}
}

