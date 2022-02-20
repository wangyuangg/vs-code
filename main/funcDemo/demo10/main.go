package main

import (
	"errors"
	"fmt"
)

func main() {
	// test03()
	// fmt.Println("main end")
	test02()
	fmt.Println("main end")
}

//函数去地区配置文件init.config信息
//自定义错误类型
func read(name string) (err error){
	if name == "config.ini" {
		return nil
	}else{
		return errors.New("读取文件错误")
	}
}
func test02(){
	err := read("config2.ini")
	if err != nil {
		fmt.Println("读取文件错误")
		panic(err)
	}
	fmt.Println("读取文件成功")
}
//默认错误类型
// func test03(){
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 		fmt.Println("err:",err)
// 		}
// 	}()	
// 	number1 :=10
// 	number2 :=0
// 	result := number1/number2
// 	fmt.Println("result:",result)
// }
