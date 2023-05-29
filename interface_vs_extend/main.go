package main

import "fmt"

type Monkey struct {
	Name string
}
type LittleMonkey struct {
	Monkey
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}
func main() {

}
