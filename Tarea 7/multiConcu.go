/*Implementa un programa en Go que simula la ejecución de múltiples tareas en paralelo utilizando goroutines.
Cada tarea realizará un trabajo específico (por ejemplo, procesamiento de datos, cálculos complejos, etc.).
Utiliza canales para comunicar el progreso de cada tarea y sincronizar la finalización de todas las goroutines.*/

package main

import (
	"fmt"
	"sync"
	"time"

)

func empleado(id int, tareas <-chan int, resultados chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range tareas {
		fmt.Printf("El empleado %d ha comenzado la tarea %d\n", id, j)
		time.Sleep(time.Second) 
		fmt.Printf("El empleado %d finalizó el trabajo %d\n", id, j)
		resultados <- j * 2
	}
}

func main() {
	numTareas := 15
	tareas := make(chan int, numTareas)
	resultados := make(chan int, numTareas)

	for j := 1; j <= numTareas; j++ {
		tareas <- j
	}
	close(tareas) //

	var wg sync.WaitGroup
	numEmpleados := 2 

	for w := 1; w <= numEmpleados; w++ {
		wg.Add(1)
		go empleado(w, tareas, resultados, &wg)
	}

	go func() {
		wg.Wait()
		close(resultados)
	}()

	for resultado := range resultados {
		fmt.Printf("Resultado: %d\n", resultado)
	}
}
