package main

import (
	"fmt"
	"os"
)

func main() {
	os.Create("./nihao.txt")

	file, err := os.Open("nihao1.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Println("file=%v", file)
}
