package main

import "fmt"

func main() {


	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20

	fmt.Println(m)

	delete(m,"c")
	fmt.Println(m["c"]) //retorn 0, não tem null

	//valor, exists
	_, exists := m["c"]

	fmt.Println(exists)//false

	value , exists := m["a"]
	fmt.Println(value)//10

	//form de criação
	var x = map[string]int{}

	x["1"] = 1
	fmt.Println(x)//map[1:1]

	var y = map[string]int{"x": 3}
	fmt.Println(y)//map[x:3]

	if v, exists := m["X"]; exists{
		fmt.Println(v)
	}else{
		fmt.Println("Não existe")
	}

	//Não existe

}

