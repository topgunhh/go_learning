package main

import "fmt"

type Person struct {
	name string
}

func main() {
	p1 := Person{
		name: "tom",
	}
	fmt.Printf("p1:%T\n", p1)
	p2 := &Person{
		name: "tom",
	}
	fmt.Printf("p2:%T\n", p2)
}
