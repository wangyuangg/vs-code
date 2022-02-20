package main

import "fmt"

func main() {
	var i integer = 1
	i.print()
	i.change()
	fmt.Println(i)
	stu := Student{Name: "Tom", Age: 18}
	fmt.Println(&stu)
}

type integer int

type Student struct {
	Name string
	Age  int
}

func (i *integer) print() {
	fmt.Println(*i)
}
func (i *integer) change() {
	*i = *i + 1
}
func (str *Student) String() string {
	return fmt.Sprintf("%s:%d", str.Name, str.Age)
}

