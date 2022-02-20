package main

import "fmt"

func main() {
	// var char byte
	// fmt.Println("请输入一个字符：")
	// fmt.Scanf("%c", &char)
	// switch char {
	// case 'a':
	// 	fmt.Println("A")
	// case 'b':
	// 	fmt.Println("B")
	// case 'c':
	// 	fmt.Println("C")
	// case 'd':
	// 	fmt.Println("D")
	// case 'e':
	// 	fmt.Println("E")
	// case 'f':99

	// 	fmt.Println("F")
	// default:
	// 	fmt.Println5("没有匹配到")

	// }
	// var score float64
	// fmt.Println("请输入一个成绩：")
	// fmt.Scanln(&score)
	// switch int(score / 60) {
	// case 1:
	// 	fmt.Println("及格")
	// case 0:
	// 	fmt.Println("不及格")
	// default:
	// 	fmt.Println("输入有误")

	// }
	var month byte
	fmt.Println("请输入一个月份：")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("输入有误")
	}

}
