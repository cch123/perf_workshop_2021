package main

func main() {
	// make slice
	// 空间开的比较大，是为了让这个 slice 分配在堆上，栈上的 slice 结果不太一样
	var sl = make([]int, 100000)
	println(sl)

	// make channel
	var ch = make(chan int, 5)
	println(ch)

	// make map
	var m = make(map[int]int, 22)
	println(m)
}
