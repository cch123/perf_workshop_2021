package main

// 这个函数只在 i == 1000 的时候会触发 bug，
// 现在的环境没法改代码，怎么通过 delve 来触发这个 bug？
func someBugFunction(i int) {
	if i == 1000 {
		panic("here bug go go")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		someBugFunction(i)
	}
}

