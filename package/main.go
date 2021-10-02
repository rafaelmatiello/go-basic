package main

import (
	"fmt"
	"package/car"
)
//go mod init -> para criar o modulo
func main() {

	c := car.Car{Nome: "BMW"}
	fmt.Println(c)
	fmt.Println(c.Inc())
	fmt.Println(c.Dec())
	fmt.Println(c.Inc())
	fmt.Println(c.Inc())
	fmt.Println(c.Inc())

}
