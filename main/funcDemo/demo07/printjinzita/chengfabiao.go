package printjinzita

import "fmt"

func Chengfabiao(totallevel int) {
	for i := 1; i <= totallevel; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d= %d ", j, i, i*j)
		}
		fmt.Println()
	}
}