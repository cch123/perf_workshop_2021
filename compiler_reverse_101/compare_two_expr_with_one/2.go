package main

type person struct {
	age int
}

func main() {
	var b = person{111}
	var a = &b
	println(a)
}

