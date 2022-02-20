package main

import "fmt"

func main() {
	var stu = Student{
		name:   "tom",
		gender: "mail",
		age:    13,
		id:     1000,
		score:  99.98,
	}
	fmt.Println(stu.say())
	var box Box = Box{
		len:   1.1,
		width: 2.0,
		heigh: 3.0,
	}
	fmt.Printf("%.2f", box.getJI())

}

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infSto := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, 
		student.age, student.id, student.score)
	return infSto
}

func (box *Box) getJI() float64 {
	return box.len * box.width * box.heigh
}

type Box struct {
	len   float64
	width float64
	heigh float64
}
