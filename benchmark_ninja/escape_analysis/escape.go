package main

import "fmt"

func returnLocalPointer() *int {
	var a = 1
	return &a
}

func localTooLarge() {
	var tl = make([]int, 10000)
	_ = tl
}

func nonConstantSlice(l int) {
	var ncs = make([]int, l)
	_ = ncs
}

func fmtSeriesFunc() {
	var i = fmt.Sprint(10)
	_ = i
}

func main() {}
