package main

import "fmt"

func main() {
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum = sum + i
	//	if sum > 20 {
	//		fmt.Println("当sum大于20时，当前数是", i)
	//		break
	//	}
	//}
	var name string
	var pwd string
	var lift int = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name == "张无忌" && pwd == "888" {
			fmt.Println("恭喜你登陆成功")
			break
		} else {
			lift--
			if lift == 0 {
				break
			}
			fmt.Printf("你还有%v次机会", lift)
		}
	}
	if lift == 0 {
		fmt.Println("机会用完，你没有登陆成功")
	}
}
