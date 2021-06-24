package main

func main() {
	for i:= 0;i< 1000;i++ {
		if i % 2 == 0 {
			iDontKnow(i)
		}
	}

}

// 我不知道这个函数什么时候会被调用到
func iDontKnow(i int) {
	println("现在我知道了")
	println("你能把调用栈打印出来吗")
}

