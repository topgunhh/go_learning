package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3.实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)

}

// 按年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}
func main() {
	// 先定义一个数组/切片
	//var intSlice = []int{0, 1, -5, 10, 90}
	//sort.Ints(intSlice)
	//fmt.Println(intSlice)

	// 对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	// 看看排序前的书序
	for _, v := range heroes {
		fmt.Println(v)
	}

	// 调用sort.sort
	sort.Sort(heroes)
	fmt.Println("排序后")
	for _, x := range heroes {
		fmt.Println(x)
	}

}
