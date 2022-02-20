package main

import "fmt"

func main() {
	var mu Method
	mu.Method1(2,2)
	(&mu).judgeNum(3)
	mu.print3(3,3,"*")
}

type Method struct {
}

func (mu Method) Method1(m int, n int) {
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (method *Method) judgeNum(num int) {
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}
func (mu *Method)print3(n int ,m int,s  string){
	for i := 0; i < n; i++ {
		for i := 0; i <= m; i++ {
			fmt.Printf(s)
		}
		fmt.Println()
	}
}
