/*
Diseña una estructura de datos en Go para representar el estado de un juego del ahorcado.
La estructura debe permitir almacenar la palabra a adivinar, el progreso de las letras adivinadas y el número de intentos restantes.
Proporciona funciones para actualizar el progreso del juego y determinar si el jugador ha ganado o perdido.
*/
package main

import (
	"fmt"
	"strings"

)

type JuegoAhorcado struct {
	palabra       string
	letras        []string
	intentos      int
	intentosMax   int
	palabraOculta string
}

func NuevoJuegoAhorcado(palabra string, intentosMax int) *JuegoAhorcado {
	palabra = strings.ToUpper(palabra)
	palabraOculta := strings.Repeat("_ ", len(palabra))
	return &JuegoAhorcado{

		palabra:       palabra,
		letras:        make([]string, 0),
		intentos:      intentosMax,
		intentosMax:   intentosMax,
		palabraOculta: palabraOculta,
	}
}

func (j *JuegoAhorcado) AdivinarLetra(letra string) {

	letra = strings.ToUpper(letra)
	if !strings.Contains(j.palabra, letra) {
		j.intentos--
	} else {
		j.letras = append(j.letras, letra)
		palabraOculta := ""
		for _, char := range j.palabra {
			if contains(j.letras, string(char)) {
				palabraOculta += string(char) + ""
			} else {
				palabraOculta += "_ "
			}
		}

		j.palabraOculta = palabraOculta
	}
}

func (j *JuegoAhorcado) HaGanado() bool {
	return j.palabraOculta == j.palabra
}

func (j *JuegoAhorcado) HaPerdido() bool {
	return j.intentos <= 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	juego := NuevoJuegoAhorcado("programacion", 5)

	for {
		fmt.Println("Intentos restantes: ", juego.intentos)
		fmt.Println("Palabra: ", juego.palabraOculta)

		var letra string
		fmt.Print("Adivina una letra: ")
		fmt.Scanln(&letra)

		juego.AdivinarLetra(letra)

		if juego.HaGanado() {
			fmt.Println("Has ganado! La palabra era: ", juego.palabra)
			break
		} else if juego.HaPerdido() {
			fmt.Println("Has perdido! La palabra era: ", juego.palabra)
			break
	}
}
}
