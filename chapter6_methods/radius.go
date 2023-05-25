package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func main() {
	var c Circle
	c.radius = 4.0
	//res := (&c).Area2()
	// go为了偷懒，在底层编译器做了优化，(&c).Area2()等价c.Area2()

	fmt.Printf("面积是%v", res)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Area2() float64 {
	// 因为c是指针，因此我们标准的访问其字段的方式是 (*c).radius
	// 同理 (*c).radius等价于 c.radius
	return math.Pi * (*c).radius * (*c).radius
}
