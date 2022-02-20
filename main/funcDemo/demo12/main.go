package main

import "fmt"

func main() {
	//创建一个byte数组存放26个，分别放置'A'~'Z'
	//使用for循环访问所有元素并打印出来
	var mychar [26]byte
	for i := 0; i < 26; i++ {
		mychar[i] = 'A' + byte(i)
	}
	for _, v := range mychar {
		fmt.Printf("%c ", v)
	}
	//请求出一个数组的最大值，并得到对应的下标
	var arr [5]int = [5]int{1, -1, 9, 90, 11}
	maxValue := arr[0]
	maxIndex := 0

	for i := 1; i < len(arr); i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
			maxIndex = i
		}

	}
	fmt.Printf("maxValue=%v maxIndex=%v\n", maxValue, maxIndex)

	var arr2 [5]int = [5]int{1, -1, 9, 90, 12}

	sum := 0

	for _, v := range arr2 {
		sum += v
	}
	//平均值
	fmt.Printf("平均值=%v\n", float64(sum)/float64(len(arr2)))
}
