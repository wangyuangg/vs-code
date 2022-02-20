package main

import "fmt"

func main() {
	var stu = Student{
		name:   "tom",
		gender: "mail",
		age:    15,
		id:     100,
		score:  99.98,
	}
	say := stu.say()
	fmt.Println(say)
	var v Visitor
	for {
		fmt.Println("请输入你的名字和年龄")
		fmt.Scanln(&v.Name)
		if v.Name == "q" {
			fmt.Println("退出")
			break
		}
		fmt.Scanln(&v.Age)
		v.showPrice()
	}

}

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}
type Visitor struct {
	Name string
	Age  int
}

func (vs *Visitor) showPrice() {
	if vs.Age >= 90 || vs.Age <= 8 {
		fmt.Println("考虑到安全 就不要玩耍")
		return
	}
	if vs.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元", vs.Name, vs.Age)
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费", vs.Name, vs.Age)
	}
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score,
	)
	return infoStr
}
