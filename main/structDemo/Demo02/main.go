package main

import "fmt"

func main() {
	var p Person
	p.Name = "张三"
	p.sum()
}

type Person struct {
	Name string
}

//func (person Person) test() {
//	fmt.Println("test", person.Name)
//}
//func (person Person) speak() {
//	fmt.Println("you are a good people")
//}
func (person *Person) sum() {
	result := 0
	for i := 0; i <= 1000; i++ {
		result += i
	}
	fmt.Println(result, (*person).Name)
}
