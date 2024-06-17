/*
	Implementar una estructura de datos para almacenar la información de cada tarea (nombre, descripción, responsable y estado).

Permitir la creación de nuevas tareas con un estado inicial de "pendiente".
Permitir la asignación de un responsable a cada tarea.
Permitir la actualización del estado de una tarea a "en progreso" o "completada".
Mostrar todas las tareas pendientes, incluyendo su nombre, descripción, responsable y estado.
*/
package main

import (
	"fmt"

)

type Tasks struct {
	ID          int
	Name        string
	Description string
	Responsible string
	Status      string
}

var listtasks []Tasks

func addTasks(id int, name string, description string, responsible string) {
	task := Tasks{
		ID:          id,
		Name:        name,
		Description: description,
		Responsible: responsible,
		Status:      "pendiente", // Status por defecto
	}
	listtasks = append(listtasks, task)
	fmt.Println("La tarea se agregó correctamente")
}

func updateTasks(id int, name string, description string, responsible string) error {
	for i := range listtasks {
		if listtasks[i].ID == id {
			listtasks[i].Name = name
			listtasks[i].Description = description
			listtasks[i].Responsible = responsible
			listtasks[i].Status = "Pendiente" // Status por defecto
			fmt.Println("La tarea se actualizó correctamente\n", listtasks[i])
			return nil
		}
	}
	return fmt.Errorf("No se encontró la tarea con ID %d", id)
}

func deleteTasks(id int) error {
	for i, tarea := range listtasks {
		if tarea.ID == id {
			listtasks = append(listtasks[:i], listtasks[i+1:]...)
			fmt.Println("La tarea se eliminó correctamente\n", tarea)
			return nil
		}
	}
	return fmt.Errorf("No se encontró la tarea con ID %d", id)
}

func listTasks() {
	fmt.Println("Inventario de tareas:\n")
	fmt.Println("ID | Name | Description | Responsible | Status")
	for _, tarea := range listtasks {
		fmt.Printf("%d | %s | %s | %s | %s\n", tarea.ID, tarea.Name, tarea.Description, tarea.Responsible, tarea.Status)
	}
	fmt.Println()
}

func main() {
	addTasks(1, "acomodar", "libros en su lugar", "Tomas")
	addTasks(2, "limpiar", "limpiar baño", "Tomas")
	addTasks(3, "Cortar pelo", "cortar pelo al perro", "Tomas")
	listTasks()
	updateTasks(1, "desacomodar", "libros desordenados", "Tomas")
	listTasks()
	deleteTasks(2)
	listTasks()
}
