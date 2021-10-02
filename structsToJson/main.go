package main

import (
	"encoding/json"
	"fmt"
)


type Car1 struct {
	Name string
	Year int `json:"-"` // quando for json, não incluir
	Color string
}

type Car2 struct {
	Name string
	Year int `json:"-"` // quando for json, não incluir
	// color não exportado
	color string
}


type Car3 struct {
	Name string `json:"Model"` //mudando o nome
	Year int `json:"-"` // quando for json, não incluir
	// color não exportado
	Color string
}



func main() {

	carGol1 := Car1{"Gol", 2019, "Amarelo"}
	carGol2 := Car2{"Gol", 2019, "Amarelo"}
	carGol3 := Car3{"Gol", 2019, "Amarelo"}

	result, _ := json.Marshal(carGol1)
	fmt.Println(string(result))//{"Name":"Gol","Color":"Amarelo"}


	result, _ = json.Marshal(carGol2)
	fmt.Println(string(result)) //{"Name":"Gol"}


	result, _ = json.Marshal(carGol3)
	fmt.Println(string(result)) //{"Model":"Gol","Color":"Amarelo"}
}