package main

import "fmt"

func main() {
	var calcuator Calcuator = Calcuator{1, 2}
	f := calcuator.jisuan("+")
	fmt.Println(f)
}
type Calcuator struct {
	num1 float64
	num2 float64
}
func (calcuator *Calcuator) jisuan(op string) float64 {
	switch op {
	case "+":
		return calcuator.num1 + calcuator.num2
	case "-":
		return calcuator.num1 - calcuator.num2
	case "*":
		return calcuator.num1 * calcuator.num2
	case "/":
		return calcuator.num1 / calcuator.num2
	default:
		return 0
	}
}