package main

import "fmt"

type Dog struct {
	name  string
	color string
	age   int
}

type Master struct {
	dog  Dog
	name string
	age  int
}

func main() {
	var tom Master
	tom.dog.color = "white"
	tom.dog.age = 2
	tom.dog.name = "little white"
	tom.age = 21
	tom.name = "topgun"

	fmt.Printf("tom:%v\n", tom)
}
