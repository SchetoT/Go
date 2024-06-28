package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	progressCh := make(chan string)

	var wg sync.WaitGroup

	wg.Add(3)

	go tareaProcesamientoDatos(ctx, progressCh, &wg)
	go tareaCalculosComplejos(ctx, progressCh, &wg)
	go tareaGeneracionInformes(ctx, progressCh, &wg)

	go func() {
		defer close(progressCh)

		for progress := range progressCh {
			fmt.Println(progress)
		}
	}()

	wg.Wait()
	fmt.Println("Todas las tareas han terminado.")
}

func tareaProcesamientoDatos(ctx context.Context, progressCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			progressCh <- fmt.Sprintf("Tarea de procesamiento de datos completada %d/%d", i, 5)
		}
	}

	progressCh <- "Tarea de procesamiento de datos completada."
}

func tareaCalculosComplejos(ctx context.Context, progressCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			progressCh <- fmt.Sprintf("Tarea de cálculos complejos completada %d/%d", i, 3)
		}
	}

	progressCh <- "Tarea de cálculos complejos completada."
}

func tareaGeneracionInformes(ctx context.Context, progressCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 2; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Duration(rand.Intn(700)) * time.Millisecond)
			progressCh <- fmt.Sprintf("Tarea de generación de informes completada %d/%d", i, 2)
		}
	}

	progressCh <- "Tarea de generación de informes completada."
}
