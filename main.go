package main

import (
	"fmt"
	"sync"
	"time"
)

//go => GoRoutine
// GoRoutine => es un hilo de ejecución ligero virtual
// Channel => para comunicar cosas entre GoRoutines

// flecha de entrada al canal, recibo del canal
func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde la go routine"
}

func imprimirMensaje(canal <-chan string) {
	msg := <-canal
	fmt.Println(msg)
}

func main() {
	canal := make(chan string)
	go decirHola(canal)
	imprimirMensaje(canal)

	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}
		close(canal2)
	}()
	//range para leer
	//canal para leer y escribir
	for num := range canal2 {
		fmt.Println("Numero recibido: ", num)
	}

	//Ejemplo mutex
	var contador int
	var mu sync.RWMutex //lector y escritor mutex
	//Mutex es para reservar

	//Writer
	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//Reader
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Printf("Leyendo desde la GoRutine %d; %d\n", id, contador)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)

}
