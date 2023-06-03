package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("nihao.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Println("file=", *file)

	defer file.Close() // 要及时关闭文件句柄，否则会有内存泄露

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF 代表文件的末尾
			break
		}
		fmt.Print(str)
	}
}
