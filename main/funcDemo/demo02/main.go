package main

import "fmt"

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}
func sum(n1 int, args ...int) int {
	sum := n1
	for _, v := range args {
		sum += v
	}
	return sum
}