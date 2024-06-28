/*
Crea un programa en Go que simula una carrera entre varios competidores utilizando goroutines.
Cada competidor debe tener una velocidad aleatoria y avanzar a intervalos de tiempo regulares.
Utiliza canales para representar la pista de carrera y coordinar el progreso de los competidores.
La Goroutine principal debe monitorear el progreso de la carrera y determinar al ganador.*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numCompetidores  = 6
	distanciaCarrera = 150
)

func main() {
	rand.Seed(time.Now().UnixNano())

	canalCarrera := make(chan int)

	var wg sync.WaitGroup
	wg.Add(numCompetidores)

	for i := 1; i <= numCompetidores; i++ {
		go competidorConPosicion(i, canalCarrera, &wg)
	}

	go func() {
		wg.Wait()
		close(canalCarrera)
	}()

	var ganador int
	var posicionGanador int = -1
	for posicion := range canalCarrera {
		if posicion > posicionGanador {
			posicionGanador = posicion
			ganador++
			fmt.Printf("Lleva la delantera el competidor %d!\n", ganador)
		}

		if posicionGanador >= distanciaCarrera {
			fmt.Printf("El competidor %d ha ganado la carrera!\n", ganador)
			break
		}
	}
}

func competidorConPosicion(id int, canalCarrera chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	velocidad := rand.Intn(8) + 2

	for posicion := 0; posicion <= distanciaCarrera; {
		posicion += velocidad
		canalCarrera <- posicion
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
