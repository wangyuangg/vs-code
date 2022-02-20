package main

import (
	"fmt"
)
func main() {
	a := make(map[string]string, 10)
	a["name"] = "tom"
	a["age"] = "18"
	fmt.Println(a)
	city :=make(map[string]string)
	city["北京"] = "beijing"
	city["上海"] = "shanghai"
	city["广州"] = "guangzhou"
	fmt.Println(city)
	var heros  = map[string]string{
		"name":"tom",
		"age":"18",
	}
	fmt.Println(heros)
	students :=make(map[string]map[string]string,10)
	students["stu01"] =make(map[string]string,3)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "男"
	students["stu01"]["address"] ="北京"
	students["stu02"] =make(map[string]string,3)
	students["stu02"]["name"] = "mary"
	students["stu02"]["sex"] = "女"
	students["stu02"]["address"] ="北京"
	fmt.Println(students)
	for _, v := range students {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}

}
