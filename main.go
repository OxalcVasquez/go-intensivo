package main

import (
	"errors"
	"fmt"
)

// funcion classic
func suma(a, b int) int {
	return a + b
}

// funcion que devuelve mas de un valor
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Errorf("No se puede devidir por 0")
		return 0, errors.New("No se puede divir por cero")
	}
	cociente := a / b
	return cociente, nil
}

// función con número variable de arguments
func imprimirNombres(nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

// ejemplo de closure
// scope chaning
var a = 0

func contador() func() int {
	c := 0
	return func() int {
		c++
		return c + a
	}
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func main() {

	cociente, error := dividir(10, 3)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(cociente)

	//Closure
	cont := contador()
	fmt.Println("contador: ,", cont())
	fmt.Println("contador: ,", cont())
	fmt.Println("contador: ,", cont())
	fmt.Println("contador: ,", cont())
	fmt.Println("contador: ,", cont())
	fmt.Println("contador: ,", cont())

	rect := Rectangulo{Ancho: 10, Alto: 5}
	fmt.Println("Area: ,", rect.Area())

}
