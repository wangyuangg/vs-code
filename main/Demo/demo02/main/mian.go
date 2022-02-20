package main

import (
	"fmt"
	"baidu.com/v/main/Demo/demo02/model"
)

func main() {
  p := model.NewPerson("Tom")
  p.SetAge(20)
  p.GetAge()
  p.SetSal(5000)
  p.GetSal()
  fmt.Println(p.Name,p.GetAge(),p.GetSal())
}