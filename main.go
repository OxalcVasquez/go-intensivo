package main

import (
	"fmt"
	"strings"
)

func main() {
	//Números
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)

	//Texto
	mensaje := "Hola, "
	concatenado := mensaje + "Oxalc"
	enMayuscula := strings.ToUpper(concatenado)

	//Booleanos
	esVerdadero := true

	// Arrays y Slice
	arrayFijo := [3]int{1, 2, 3}
	sliceVariable := []int{4, 5, 6}
	sliceVariable = append(sliceVariable, 7)

	//Mapas - Diccionarios
	diccionario := map[string]int{
		"clave1": 1,
		"clave2": 2,
	}

	//Structs, es un tipo
	type Persona struct {
		Nombre string
		Edad   int
	}

	persona := Persona{Nombre: "Oxalc", Edad: 30}

	//Imprimer Resultados
	fmt.Println("Suma", suma)
	fmt.Println("Mensaje", enMayuscula)
	fmt.Println("Array", arrayFijo)
	fmt.Println("Slice", sliceVariable)
	fmt.Println("Mapa", diccionario)
	fmt.Println("Struct", persona)
	fmt.Println("Bool", esVerdadero)

}
