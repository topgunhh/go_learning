package main

import (
	"fmt"
	"go_learning/chapter11_encapsulation/model"
)

func main() {
	p := model.NewPerson("tom")
	p.SetAge(18)
	p.SetSal(15000)
	fmt.Println(p.Name, "salary is", p.GetSal(), "age is", p.GetAge())
	fmt.Println(*p)

}
