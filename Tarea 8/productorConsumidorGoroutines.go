/*
Desarrolla un programa en Go que implemente el patrón productor-consumidor utilizando goroutines y canales.
El productor debe generar datos aleatorios y enviarlos a un canal, mientras que uno o varios consumidores deben recibir los datos del canal y
procesarlos de alguna manera (por ejemplo, calcular estadísticas, almacenar en una base de datos, etc.).
Asegúrate de implementar la sincronización adecuada para evitar condiciones de carrera.*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

)

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)

	go productor(ch, done)

	go consumidor(ch, &wg)
	go consumidor(ch, &wg)

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
	fmt.Println("Todos los consumidores han terminado.")
}

func productor(ch chan<- int, done chan<- bool) {
	defer close(ch)

	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(100)
		fmt.Printf("Productor produce: %d\n", randomNumber)
		ch <- randomNumber
		time.Sleep(time.Millisecond * 500)
	}

	done <- true
}

func consumidor(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for number := range ch {
		fmt.Printf("Consumidor recibe: %d\n", number)
		time.Sleep(time.Second)
	}
	fmt.Println("Consumidor ha terminado.")
}
