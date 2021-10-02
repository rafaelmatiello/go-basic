package main

import "fmt"

func main() {

	//make cria em memória
	// tipo, tamanho, capacidade.
	//capacidade quando faz um append, ele adicionar mais 5 posições, vai 5 para 10, mais 20, mais 40, sempre dobra
	slice := make([]int, 2, 5)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	//adicionar dinamicamente.
	slice = append(slice, 1, 10, 293,4)
	fmt.Println(slice)
	fmt.Println(len(slice))


	sliceX := make([]int, 2,5)
	fmt.Println(sliceX)
	sliceY := sliceX;
	sliceX[0] = 9

	fmt.Println(sliceX)
	fmt.Println(sliceY)
	//[9 0]
	//[9 0]

	//depois de estourar perde a referência
	sliceY = append(sliceY, 1, 10, 293,4)
	fmt.Println(sliceX)
	fmt.Println(sliceY)

	//[9 0]
	//[9 0 1 10 293 4]

	//ja estourou então vai perder a referência
	sliceX[0] = 8

	fmt.Println(sliceX)
	fmt.Println(sliceY)
	//[8 0]
	//[9 0 1 10 293 4]


	fmt.Println("--------------")

	//inicialização
	sliceR := []string{
		"Hello",
		"World",
	"Rafael"}

	fmt.Println(sliceR[0]) //Hello
	//do inicio até o 2
	fmt.Println(sliceR[:2]) //[Hello World]
	fmt.Println(sliceR[1:2]) // World
	// do indice ao final
	fmt.Println(sliceR[1:]) // [World Rafael]



}

