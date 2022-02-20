package main

import "fmt"

func main() {
	//统计三个班成绩情况，每个班有5个同学，求出各个班的平均分和所有班的平均分【数据从键盘输出】
	var classNum int = 3
	var stuNum int = 5
	var totalsun float64
	for j := 1; j <= classNum; j++ {
		sun := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班第%d个学生的成绩\n", j, i)
			fmt.Scanln(&score)
			sun += score
		}
		fmt.Printf("第%d班级的平均分是%v\n", j, sun/float64(stuNum))
		totalsun += sun

	}
	fmt.Printf("所有班级的平均分是%v\n", totalsun / (float64(classNum)*float64(stuNum)))
}
