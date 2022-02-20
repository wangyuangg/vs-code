package main

import "fmt"

func main() {
	//使用while方式输出10据""hello,world"
	var i int = 1
	for  {
		if i > 10 {
			break
		}
		fmt.Println("helloworld",i)
		i++
	}
		//使用do-while方式输出10句""hello,world"
	var j int = 1
	for {
		fmt.Println("hello,worldok",j)
		j++
		if j > 10 {
			break//跳出for循环
		}
		
	}
}