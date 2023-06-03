package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	// 账户余额
	balance := 10000.0
	// 收支金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 记录是否有支出
	flag := false
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对detail进行拼接处理即可
	detail := "收支\t账户金额\t收支金额\t说   明"

	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记收入")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出")
		fmt.Println("")
		fmt.Println("请选择（1-4）:")

		fmt.Scanln(&key)

		switch key {
		case "1":
			if flag {
				fmt.Println("----------当前收支明细记录----------")
				fmt.Println(detail)
			}
			fmt.Println("当前没有收支明细")
		case "2":
			fmt.Scanln(&money)
			flag = true
			fmt.Printf("本次收入金额：%v\n", money)
			balance = balance + money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)

			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			fmt.Println(detail)
		case "3":
			fmt.Println("登记支出..")
			fmt.Scanln(&money)
			flag = true
			if money > balance {
				fmt.Println("余额的金额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			fmt.Println(detail)
		case "4":
			fmt.Println("请问是否要退出？y/n")
			judge := ""
			for {
				fmt.Scanln(&judge)
				if judge == "y" || judge == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入 y/n")

			}
			loop = false
		default:
			fmt.Println("请输入正确的选项")
		}

		if !loop {

			break
		}

	}
	fmt.Println("已退出软件")
}
