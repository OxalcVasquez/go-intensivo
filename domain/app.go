package domain

import (
	"fmt"
	"learning-go/domain/entity"
)

func App() {
	persona := entity.Persona{Nombre: "Jordan", Apellido: "Vasquez", Edad: 32}

	fmt.Println(persona.Saludar())
	persona.CumplirAnios()

	fmt.Println(persona)
}
