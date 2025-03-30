package main

import "fmt"

var a = 1

// referencia de nombre "a" que apunta a un espacio de memoria que contiene en su interior el valor 1

func incrementar(numero *int) {
	//NUNCA modifiques el argumento directamente
	*numero++
}

// funcion que devuelve mas de un valor
func main() {

	valor := 10
	fmt.Println("Valor antes de incrementar:", valor)
	incrementar(&valor) //dirección de memoria
	fmt.Println("Valor despues de incrementar:", valor)

	// new()
	puntero := new(int) //puntero int inicializado en 0, donde se encuentra algo 0xc00000a0f8
	fmt.Println("Valor inicial con new: ", puntero)
	*puntero = 20
	fmt.Println("Valor inicial con new: ", *puntero)

}
