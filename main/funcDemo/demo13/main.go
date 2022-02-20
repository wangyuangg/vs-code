//1.随机生成5个数
package main

import (
	"fmt"
)

func main() {
	// var intArr3 [5]int
	// rand.Seed(time.Now().Unix())
	// for i := 0; i < len(intArr3); i++ {
	//intArr3[i] = rand.Intn(100)
	// }
	// fmt.Println("交换前=", intArr3)

	// temp1 := 0
	// for i := 0; i < len(intArr3) / 2; i++ {
	// 	temp1 = intArr3[len(intArr3)-1-i]
	// 	intArr3[len(intArr3)-1-i] = intArr3[i]
	// 	intArr3[i] = temp1
	// }
	// fmt.Println("交换后=", intArr3)
	//var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	//
	//slice := intArr[1:3] //切片intArr[]中的第二个元素开始，到第三个元素结束不包含第三个元素
	//fmt.Println("intArr=", intArr)
	//fmt.Println("slice=", slice)
	//fmt.Println("len(slice)=", len(slice))
	//fmt.Println("slice 的容量=", cap(slice)) //切片的容量是可以动态变化的
	//fmt.Printf("intArr[1]的地址值=%p\n", &intArr[1])
	//fmt.Printf("slice[0]的地址值=%p\n", &slice[0])
	fmt.Println("--------------------------")
	var slice = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("--------------------------")
	var strSlice = []string{"tom", "jack", "mary"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}
	for _, value := range strSlice {
		fmt.Println(value)
	}
	var slice3 = []int{100, 200, 300}
	fmt.Println("slice3=", slice3)
	ints := append(slice3, 400, 500, 600)
	fmt.Println(ints)
	var a = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, 10)
	fmt.Println(slice4)
	copy(slice4, a)
	fmt.Println(slice4)
}
