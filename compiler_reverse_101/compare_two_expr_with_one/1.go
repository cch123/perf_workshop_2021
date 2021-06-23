package main

type person struct {
	age int
}

func main() {
	var a = &person{111}
	println(a)
}
