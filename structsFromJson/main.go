


package main

import (
"encoding/json"
"fmt"
)


type Car struct {
	Name string
	Year int
	Color string
}



func main() {

	var car Car

	data := []byte(`{"Name": "Gol", "Year": 2010, "Color": "Verde"}`)

	//lembrar de passar como referência
	json.Unmarshal(data, &car)

	fmt.Println(car.Name, car.Year, car.Color)


}

