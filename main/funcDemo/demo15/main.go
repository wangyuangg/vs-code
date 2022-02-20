package main

import "fmt"

func main() {
	//【24，69，80，57，13】
	var intArr = [5]int{24, 69, 80, 57, 13}
	BubbleSort(intArr)
	//fmt.Println("排序后", intArr)
}
func BubbleSort(arr [5]int) {
	fmt.Println("排序前", arr)
	temp := 0
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}

	}
	fmt.Println("排序后", arr)
}
