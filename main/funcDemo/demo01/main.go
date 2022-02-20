package main

import "fmt"

func main() {
	//编写一个函数test
	n1 := 4
	test(n1)
}

func test(n1 int) {
	if n1 > 2 {
		n1--
		test(n1)
	}
	fmt.Println(n1)
}