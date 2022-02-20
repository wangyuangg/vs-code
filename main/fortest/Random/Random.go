package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			fmt.Println(n)
			break
		}
	}
	fmt.Println("count:", count)
}
