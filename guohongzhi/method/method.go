package main

import "fmt"

type Customer struct {
	name string
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {
	cus := Customer{
		name: "tom",
	}

	b := cus.login("tom", "123")
	fmt.Printf("b:%v\n", b)
}
