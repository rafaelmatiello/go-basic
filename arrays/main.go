package main

import "fmt"

func main() {

	var x [10]int
	//[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(x)
	fmt.Println(len(x))

	x[0] = 1
	x[1] = 6
	fmt.Println(x)

	z := [5]int{1,3,54,5,6}
	fmt.Println(z)
}

