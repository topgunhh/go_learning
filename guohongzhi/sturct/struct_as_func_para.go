package main

import "fmt"

type Person struct {
	id   int
	name string
}

// 值传递拷贝了一份副本
func showPerson(person Person) {
	person.id = 1
	person.name = "kite"
	fmt.Printf("person:%v\n", person)
}

func showPersonPoint(person2 *Person) {
	person2.id = 1
	person2.name = "kite"
	fmt.Printf("person:%v\n", person2)

}

func main() {
	person := Person{1, "tom"}
	fmt.Printf("person:%v\n", person)
	fmt.Println("--------------------")
	showPerson(person)
	fmt.Println("--------------------")
	fmt.Printf("person:%v\n", person)

	fmt.Println("--------------------")
	fmt.Println("传递结构体指针")
	person2 := Person{2, "lite"}
	fmt.Printf("person2:%v\n", person2)
	fmt.Println("--------------------")
	showPersonPoint(&person2)

	fmt.Println("--------------------")
	fmt.Printf("person2:%v\n", person2)

}
