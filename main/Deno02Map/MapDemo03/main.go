package main

import "fmt"

func main() {
	//map切片的使用、
	var monstres []map[string]string
	monstres = make([]map[string]string, 2)
	if monstres[0] == nil {
		monstres[0] = make(map[string]string, 2)
		monstres[0]["name"] = "牛魔王"
		monstres[0]["age"] = "500"
	}
	if monstres[1] == nil {
		monstres[1] = make(map[string]string, 2)
		monstres[1]["name"] = "玉兔"
		monstres[1]["age"] = "700"
	}
	newMonstres := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}
	monstres = append(monstres, newMonstres)
	fmt.Println(monstres)
}
