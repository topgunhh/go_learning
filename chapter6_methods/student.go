package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func main() {
	var stu = Student{
		name:   "",
		gender: "",
		age:    0,
		id:     0,
		score:  0,
	}
	stu.say()

}

func (s *Student) say() string {

	infoStr, _ := fmt.Printf("%v,%v,%v,%v,%v", s.name, s.id, s.gender, s.age, s.score)
	return string(infoStr)
}
