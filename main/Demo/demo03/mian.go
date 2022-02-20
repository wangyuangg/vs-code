package main

import "fmt"

func main() {
	var b B
	b.Name = "tom"
	b.hello()
	b.SayOK()
	
}

type A struct {
	Name string
	//age  int
}

func (a *A) SayOK() {
	println("SayOK", a.Name)
}

func (a *A) hello() {
	fmt.Println("hello", a.Name)
}

type B struct {
	A
}
