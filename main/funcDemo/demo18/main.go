package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	//fmt.Println(arr)
	for i :=0;i < len(arr);i++{
		for j := 0;j < len(arr[i]);j++{
			fmt.Printf("%d ",arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-----------------")
	for i,v := range arr{
		for j,v1 := range v{
			fmt.Printf("%d %d %d\n",i,j,v1)
		}
	}
}
