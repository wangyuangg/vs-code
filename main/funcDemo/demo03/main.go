package main

func main() {
	n1 := 1
	n2 := 2
	swap(&n1, &n2)
	println(n1, n2)

}
//swap n1 and n2
func swap(n1, n2 *int) {
	t := *n1
	*n1=*n2
	*n2=t
}


