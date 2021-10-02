package main

import "fmt"

type CarName string
type CarYear int
type Color string

type Car struct {
	Name string
	Year int
	Color string
}

type SuperCar struct {
	//fazendo herança
	Car
	CanFly bool
}

type SuperCarName struct {
	//fazendo herança
	Car
	Name string
	CanFly bool
}



//função info dentro o objeto carro
func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s\n", c.Name, c.Year, c.Color)
}

func main() {

	car1 := Car{"Corola", 2019, "white"}
	car2 := Car{"BMW 320", 2012, "red"}



	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car2.Name, car2.Year, car2.Color)

	fmt.Println(car1.info())

	//{Corola 2019 white}
	//{BMW 320 2012 re}
	//BMW 320 2012 re
	//Car: Corola
	//Year: 2019
	//Color: white

	carSuper := SuperCar{Car: Car{"BMW 550", 2013, "blue"}, CanFly: true}
	fmt.Println(carSuper) //{{BMW 550 2013 blue} true}

	fmt.Println(carSuper.Name) // BMW 550
	fmt.Println(carSuper.info())
	//Car: BMW 550
	//Year: 2013
	//Color: blue


	superCarName := SuperCarName{Car: car2, CanFly: true, Name: "Fox"}

	fmt.Println(superCarName) //{{BMW 320 2012 red} Fox true}
	fmt.Println(superCarName.Name) // Fox
	fmt.Println(superCarName.Car.Name) // BMW 320
}

