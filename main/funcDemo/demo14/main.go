package main

import "fmt"

func main() {
	//str := "hello,world!"
	//slice := str[6:]
	//fmt.Println(slice)
	//arr1 := []rune(str)
	//fmt.Println(arr1, "**")
	//arr1[0] = '北'
	//str = string(arr1)
	//fmt.Println(str)
	fmt.Println(fbn(10))
}
func fbn(n int) []uint64 {
	slice := make([]uint64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
