package main

import (
	"fmt"
	"package/car"
)

func main() {

	c := car.Car{Nome: "BMW"}
	fmt.Println(c)
	fmt.Println(c.Inc())
	fmt.Println(c.Dec())
	fmt.Println(c.Inc())
	fmt.Println(c.Inc())
	fmt.Println(c.Inc())

}
