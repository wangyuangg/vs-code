package main

import "fmt"

func main() {
	var stu Stu
	var a AInterface = stu
	a.say()
	var i integer = 10
	var b AInterface = i
	b.say()
}

type AInterface interface {
	say()
}
type integer int

func (i integer) say() {
	fmt.Println("integer")
}

type Stu struct {
	Name string
}

func (s Stu) say() {
	fmt.Println("hello")
}

type BInterface interface {
	hello()
}
type Monster struct {


}
func (m Monster) hello() {
	fmt.Println("hello1")
}
func (m Monster) say() {
	fmt.Println("hello2")
}