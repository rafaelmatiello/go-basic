package main

import "fmt"

func main() {

	fmt.Println(sum(10, 8))

	fmt.Println(nameReturn("Gabi"))
	x, y := moreReturns("R", "G")
	fmt.Println(nameReturn(x))
	fmt.Println(nameReturn(y))

	fmt.Println(variadicFunc(1,233,43))

	// função anônima
	z := 0
	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())

	fmt.Println(funcInsideFunc(1)())

}


func funcInsideFunc( a int) func() int{

	return func() int{
		return a + 10
	}

}

func sum(a int, b int) int{
	return a + b
}


// o retorno da função é o valor de x
func nameReturn(a string)( x string){
	x = a
	return
}

//retornar mais de um valor
func moreReturns(a string, b string) (string, string){
	return a, b
}


func variadicFunc(x ...int ) int{
	var res int
	// _ é o blank, porque todas as variáveis devem ser utilizadas
	// for indice, value := range x { // o primeiro parâmetro é o indice.
	for _, value := range x {
		res += value

	}

	return res
}


