package main

import "fmt"

//import "fmt"

func main() {
	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个学生的成绩:\n", i+1)
	// 	fmt.Scanln(&score[i])
	// }
	// for _, v := range score {
	// 	fmt.Printf("score=%v\n", v)
	// }
	// //四种初始化数组的方式
	//1.默认初始化
	var arr3 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr3=%v\n", arr3)
	//4.指定初始值
	var arr4 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr4=%v\n", arr4)
	var arr1 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr1=%v\n", arr1)
	var arr2 = [3]int{1:1,0: 2,2:3}
	fmt.Printf("arr2=%v\n", arr2)


}
