package main

import "fmt"

// 声明一个接口
type Usb interface {
	Insert()
	UnInsert()
}
type Phone struct {
}

type Laptop struct {
}

type Camera struct {
}

// 让Phone实现Usb接口的方法
func (p Phone) Insert() {
	fmt.Println("phone插入")
}
func (p Phone) UnInsert() {
	fmt.Println("phone拔出")
}
func (c Camera) Insert() {
	fmt.Println("Camera插入")
}
func (c Camera) UnInsert() {
	fmt.Println("Camera拔出")
}

// 编写一个方法 Working，接收一个Usb接口类型变量
// 只要实现了Usb接口（所谓实现接口，就是指实现了Usb接口声明的所有方法）
func (l Laptop) Working(usb Usb) {
	usb.Insert()
	usb.UnInsert()
}
func main() {
	//先创建结构体变量
	laptop := Laptop{}
	phone := Phone{}
	camera := Camera{}
	// 关键点
	laptop.Working(phone)
	laptop.Working(camera)
}
