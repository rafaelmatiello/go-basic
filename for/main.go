package main

import "fmt"

func main() {

	for i:=0; i< 5; i++{
		fmt.Println(i)
	}

	fmt.Println("------ while -------")
	x := 0
	for x < 5{
		fmt.Println(x)
		x++
	}

	fmt.Println("------ loop finito -------")
	for  {
		fmt.Println(x)
		x++

		if x > 10 {
			break
		}
	}
}

