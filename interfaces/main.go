package main

import "fmt"

type Car struct {
	Name string
}

type Motorcycle struct {
	Name string
}

// não precisa dar um implements, somente precisa ter o método
type Veiculo interface {
	start() string
}

func (c Car) start() string {
	return "Iniciando o carro: " + c.Name
}

func (m Motorcycle) start() string {
	return "Iniciando a moto: " + m.Name
}

func startVeiculo(v Veiculo) string {
	return v.start()
}

func main() {
	car := Car{"Gol"}
	moto := Motorcycle{
		Name: "Honda"}

	fmt.Println(startVeiculo(car))
	fmt.Println(startVeiculo(moto))
}
