package main

import "fmt"

func main() {
	var a Account = Account{
		Account: "工商银行",
		Pwd:     "11111111",
		Balance: 100.00,
	}
	(&a).Deposit(12345, "11111111")
	a.Query("11111111")
	(&a).WithDraw(100.00, "11111111")
	a.Query("11111111")

}

type Account struct {
	Account string
	Pwd     string
	Balance float64
}

func (account *Account) Deposit(money float64, pwd string) { //存款
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("存款金额必须大于0")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}
func (account *Account) WithDraw(money float64, pwd string) { //取款
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 || money >= account.Balance {
		{
			fmt.Println("取款金额必须大于0且小于等于余额")
			return
		}

		account.Balance -= money
		fmt.Println("存款成功")
	}
}
func (account *Account) Query(pwd string) { //查询余额
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	fmt.Println("余额为：", account.Balance)
}
