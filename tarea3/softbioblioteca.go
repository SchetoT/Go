/*
Implementar una estructura de datos para almacenar la información de cada libro (título, autor, género y estado).
Permitir la adición de nuevos libros a la colección.
Permitir la búsqueda de libros por título o autor.
Permitir la actualización del estado de un libro a "disponible" o "prestado".
Permitir la eliminación de libros de la colección.
*/

package main

import (
	"fmt"
	"strings"
)

type Libro struct {
	ID     int
	Titulo string
	Autor  string
	Genero string
	Estado string
}

var biblioteca []Libro

func addLibro(titulo string, autor string, genero string, estado string) {
	id := len(biblioteca) + 1
	libro := Libro{ID: id, Titulo: titulo, Autor: autor, Genero: genero, Estado: estado}
	biblioteca = append(biblioteca, libro)
	fmt.Println("Se agregó correctamente el libro", libro)
}

func updateLibro(id int, titulo string, autor string, genero string, estado string) error {
	for i := range biblioteca {
		if biblioteca[i].ID == id {
			biblioteca[i].Titulo = titulo
			biblioteca[i].Autor = autor
			biblioteca[i].Genero = genero
			biblioteca[i].Estado = estado
			return nil
		}
	}
	return fmt.Errorf("Libro no encontrado")
}

func deleteLibro(id int) error {
	for i := range biblioteca {
		if biblioteca[i].ID == id {
			biblioteca = append(biblioteca[:i], biblioteca[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Libro no encontrado")
}
func buscarAutor(autor string) []Libro {
	var resultados []Libro
	for _, libro := range biblioteca {
		if strings.Contains(strings.ToLower(libro.Autor), strings.ToLower(autor)) {
			resultados = append(resultados, libro)
		}
	}
	return resultados
}
func buscarTitulo(titulo string) []Libro {
	var resultados []Libro
	for _, libro := range biblioteca {
		if strings.Contains(strings.ToLower(libro.Titulo), strings.ToLower(titulo)) {
			resultados = append(resultados, libro)
		}
	}
	return resultados
}

func listLibros() {
	fmt.Println("Lista de libros: \n")
	for _, libro := range biblioteca {
		fmt.Println("ID | TITULO | GENERO | ESTADO | AUTOR \n", libro.ID, "|", libro.Titulo, "|", libro.Genero, "|", libro.Estado, "|", libro.Autor)
	}
	fmt.Println("\n")
	//Imprimo la lista de libros completa
}
func eliminarTodosLosLibros() {
	biblioteca = []Libro{}
	fmt.Println("Se eliminaron todos los libros")
	//asigno una slice "vacio" para eliminar todos los libros
}
func main() {
	addLibro("La papa", "Armando Muros", "fantasia", "prestado")
	addLibro("El melon", "Ricardo Arjona", "Terror", "reservado")
	addLibro("La sandia", "Bob Esponja", "Horror", "prestado")
	listLibros()
	//Agregando libros

	fmt.Println("Resultados de la busqueda por EL:")
	resultadosTitulo := buscarTitulo("El")
	for _, libro := range resultadosTitulo {
		fmt.Println("ID | TITULO | GENERO | ESTADO | AUTOR \n", libro.ID, "|", libro.Titulo, "|", libro.Genero, "|", libro.Estado, "|", libro.Autor)

	}
	fmt.Println()
	//Buscando libros por "El"

	fmt.Println("Resultados de la busqueda por Armando:")
	resultadosAutor := buscarAutor("Armando")
	for _, libro := range resultadosAutor {
		fmt.Println("ID | TITULO | GENERO | ESTADO | AUTOR \n", libro.ID, "|", libro.Titulo, "|", libro.Genero, "|", libro.Estado, "|", libro.Autor)

	}
	fmt.Println()
	//Buscando libros por "Armando"

	if err := updateLibro(3, "La sandia", "Bob esponja", "Horror", "reservado"); err != nil {
		fmt.Println("Se actualizó el libro!")
	}
	listLibros()
	//Actualizando el libro de bob esponja a "reservado"

	if err := deleteLibro(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se eliminó el libro!")
	}
	listLibros()
	//Eliminando libro en este caso el libro con id 2
	eliminarTodosLosLibros()
	listLibros()
	//Elimino todos los libros e imprimo el resultado, llamo la funcion al ultimo ya que no tendria sentido hacerlo antes
}
