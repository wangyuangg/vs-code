package main

import "fmt"

func main() {
	name := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)
	for i := 0; i < len(name); i++ {
		if heroName == name[i] {
			fmt.Println("找到了")
		} else if i == (len(name) - 1) {
			fmt.Println("没找到")
		}
	}
}
