package main

import "fmt"

type Names []interface{}

// detalhes, conteúdo do ponteiro, para dentro do conteúdo do poteiro
func (names *Names) load(){
	// como é uma interface vazia aceita tudo
	*names = Names{
		"Rafael",
		"Gabriela",
		1,
		true}
}

func (names Names) print(){
	fmt.Println(names)
}



func main() {

	var names Names
	names.load()
	names.print()


}
