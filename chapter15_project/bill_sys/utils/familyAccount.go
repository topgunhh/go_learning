package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//声明必要的字段
	key  string
	loop bool
	// 账户余额
	balance float64
	// 收支金额
	money float64
	// 每次收支的说明
	note string
	// 记录是否有支出
	flag bool
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对detail进行拼接处理即可
	detail string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    false,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		detail:  "",
	}
}
func (this *FamilyAccount) incomeRecord() {
	fmt.Scanln(&this.money)
	this.flag = true
	fmt.Printf("本次收入金额：%v\n", this.money)
	this.balance = this.balance + this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)

	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	fmt.Println(this.detail)
}

func (this *FamilyAccount) showDetials() {
	if this.flag {
		fmt.Println("----------当前收支明细记录----------")
		fmt.Println(this.detail)
	} else {
		fmt.Println("当前没有收支明细")
	}

}

func (this *FamilyAccount) outcomeRecord() {
	fmt.Println("登记支出..")
	fmt.Scanln(&this.money)
	this.flag = true
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	fmt.Println(this.detail)
}
func (this *FamilyAccount) exit() {
	fmt.Println("请问是否要退出？y/n")
	judge := ""
	for {
		fmt.Scanln(&judge)
		if judge == "y" || judge == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")

	}
	this.loop = false
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记收入")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出")
		fmt.Println("")
		fmt.Println("请选择（1-4）:")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetials()
		case "2":
			this.incomeRecord()
		case "3":
			this.outcomeRecord()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
			if !this.loop {
				break
			}
			fmt.Println("已退出软件")
		}
	}
}
