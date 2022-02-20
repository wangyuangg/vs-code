package main

import "fmt"

func main() {
	var totallevel int = 8
	//打印空心金字塔
	for i := 1; i <= totallevel; i++ {
		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totallevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")

			}

		}
		fmt.Println()

	}

}
