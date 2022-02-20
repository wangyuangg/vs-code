package main

import (
	"fmt"
	"sort"
)
type Stu struct {
	Name string
	Age int
	Address string
}
func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	map1[2] = 50
	fmt.Println(map1)

	var keys []int
	for k := range map1 {
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Println(map1[v])
	}
	
	student :=make(map[string]Stu,10)
	Stu1 := Stu{"tom",18,"shanghai"}
	Stu2 := Stu{"mary",28,"beijing"}
	student["no1"] = Stu1
	student["no2"] = Stu2
	fmt.Println(student)
	for k,v := range student {
		fmt.Printf("学生的编号是%s,姓名是%s,年龄是%d,地址是%s\n",k,v.Name,v.Age,v.Address)
	}
	
}
