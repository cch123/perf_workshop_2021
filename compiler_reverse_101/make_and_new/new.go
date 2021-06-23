package main

type person struct{ age int }

func main() {
	var a = new(int)
	var b = new(person)
	var c = new(chan int)
	go func() {
		x := <-*c
		println(x)
	}()

	var d = new(map[int]int)

	println(a, b, c, d)
}
