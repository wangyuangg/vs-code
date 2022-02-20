package main

import "fmt"

func main() {
	city := make(map[string]string)
	city["北京"] = "beijing"
	city["上海"] = "shanghai"
	city["广州"] = "guangzhou"
	fmt.Println(city)
	city["北京"] = "beijing1**"
	fmt.Println(city)
	//delete(city, "北京")
	fmt.Println(city)
	//val, findRes := city["北京"]
	// if findRes {
	// 	fmt.Println(val)
	// } else {
	// 	fmt.Println("没有北京")
	// }
	for _, v := range city {
		fmt.Println(v)
	}
		
	}
