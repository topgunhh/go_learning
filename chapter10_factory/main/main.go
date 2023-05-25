package main

import (
	"fmt"
	"go_learning/chapter10_factory/module"
)

func main() {
	//var stu = module.Student{
	//	Name:  "tom",
	//	Score: 100,
	//}
	//fmt.Println(stu)
	var stu = module.NewStudent("tom", 88.8)
	fmt.Println(*stu)

	fmt.Println("name=", stu.Name, "Score=", stu.Score)
}
