package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

type Pupil struct {
	Student // 嵌入了Student的匿名结构体
}

type Graduate struct {
	Student
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	stu.Score = score
}

func (p *Pupil) Testing() {

}

func (p *Graduate) Testing() {

}

func main() {
	// 当我们对结构体嵌入了匿名结构体，使用的方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 15
	pupil.Testing()
	pupil.Student.Score = 80
	pupil.Student.ShowInfo()
}
