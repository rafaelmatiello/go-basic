package car

type Car struct {
	Nome     string
	// Não será visivel por fora
	contador int
}

func (car *Car) Inc() int {
	car.contador = car.contador + 1
	return car.contador
}

func (car *Car) Dec() int {
	car.contador = car.contador - 1
	return car.contador
}
