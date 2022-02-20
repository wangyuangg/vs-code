package utils

import "fmt"

type Family struct {
	key string
	//控制循环
	details string
	balance float64
	money   float64
	note    string
	loop    bool
	flag    bool
}

//绑定相应的方法
func (f *Family) ShowDetails() {
	if f.flag {
		fmt.Println("*************当前收支明细***************")
		fmt.Println(f.details)
	} else {
		fmt.Println("没有收支明细")
	}

}
func (f *Family) income() {
	fmt.Println("请输入收入金额")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("请输入收入说明")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", "账户", f.money, f.note)
	f.flag = true
}
func (f *Family) pay() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足")
		//return
	}
	f.balance -= f.money
	fmt.Println("请输入支出说明")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", "账户", f.money, f.note)
	f.flag = true
}
func (f *Family) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入有误，请重新输入")
	}
	if choice == "y" {
		f.loop = false
	}

}
func (f *Family) ShowMenu() {

	for {
		fmt.Println("*************家庭收支软件***************")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Println("请选择1~4")
		fmt.Scanln(&f.key)
		switch f.key {
		case "1":
			f.ShowDetails()

		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !f.loop {
			break
		}
	}
}

func NewFamily() *Family {
	return &Family{
		key:     "",
		details: "收支\t账户\t金额\t说  明\n",
		balance: 10000.0,
		money:   0.0,
		note:    "",
		loop:    true,
		flag:    false,
	}
}
