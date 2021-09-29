package main

import "fmt"


func main() {
	x := 10

	y := &x

	fmt.Println(y)
	fmt.Println(*y)

	*y = 20

	fmt.Println(y)
	fmt.Println(*y)

	var z *int = y

	fmt.Println(z)
	fmt.Println(*z)


	var w *int = &x

	fmt.Println(w)
	fmt.Println(*w)

	a := 100
	fmt.Println(sum(a))
	//101
	fmt.Println(a)
	//100

	fmt.Println(sumReference(&a))
	//101
	fmt.Println(a)
	//101
}

// passagem por cópia
func sum(a int) int{
	a = a + 1
	return a
}

//passagem por referência
func sumReference(a *int) int{
	*a = *a + 1
	return *a
}