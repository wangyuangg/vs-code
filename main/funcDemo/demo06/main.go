package main

import "fmt"

func main() {
	res := sum(1, 2)
	fmt.Println("res=", res)
}
func sum(n1, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	result := n1 + n2
	fmt.Println("ok3 result=", result)
	return result
	
}
