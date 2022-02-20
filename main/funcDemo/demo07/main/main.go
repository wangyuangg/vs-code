package main

import (
	"fmt"

	"baidu.com/v/main/funcDemo/demo07/printjinzita"
)
func main() {
	//从终端输入一个int类型的数据
	var totallevel int
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&totallevel)
	printjinzita.Jinzita(totallevel)
	printjinzita.Chengfabiao(totallevel)


}


