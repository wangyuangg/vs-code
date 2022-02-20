package main

import (
	"fmt"
	"baidu.com/v/main/structDemo/demo08/model"
)
func main() {
	// var stu = model.student{
	// 	Name:"tom",
	// 	Score:99.98,
	// }
	var stu = model.NewStudent("tom",99.98)
	
	
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"score=",stu.ReturnInt())
}
