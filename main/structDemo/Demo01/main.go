package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	monster := Monster{"牛魔王",500,"芭蕉扇·"}
	jsonStr,err :=json.Marshal(monster)
	if err != nil {
		fmt.Println("json.Marshal failed,err:",err)
	}
	fmt.Println("jsonStr:",string(jsonStr))
}
type Monster struct {
	Name string `json:"name"`//结构体的字段名要和json的key一致
	Age int `json:"age"`//结构体
	AttackTool string `json:"attackTool"`
}
