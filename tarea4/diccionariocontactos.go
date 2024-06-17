/*
Un usuario necesita un programa para gestionar sus contactos. El programa debe permitir la adición de nuevos contactos, la búsqueda de contactos por nombre o número de teléfono, la actualización de la información de un contacto y la eliminación de contactos.

Requisitos:

Implementar una estructura de datos para almacenar la información de cada contacto (nombre, número de teléfono, correo electrónico y dirección).
Permitir la adición de nuevos contactos con su información completa.
Permitir la búsqueda de contactos por nombre o número de teléfono.
Permitir la actualización de la información de un contacto (teléfono, correo electrónico o dirección).
Permitir la eliminación de contactos de la lista.
*/
package main

import (
	"fmt"
	"strings"
)

type Contacto struct {
	ID        int
	Nombre    string
	Telefono  int
	Correo    string
	Direccion string
}

var agenda []Contacto

func addContacto(nombre string, telefono int64, correo string, direccion string) {
	id := len(agenda) + 1
	contacto := Contacto{ID: id, Nombre: nombre, Telefono: int(telefono), Correo: correo, Direccion: direccion}
	agenda = append(agenda, contacto)
	fmt.Println("El contacto se agregó correctamente a la agenda\n", contacto)
}

func updateContacto(id int, nombre string, telefono int, correo string, direccion string) {
	for i := range agenda {
		if agenda[i].ID == id {
			agenda[i].Nombre = nombre
			agenda[i].Telefono = telefono
			agenda[i].Correo = correo
			agenda[i].Direccion = direccion
			fmt.Println("El contacto se actualizó correctamente\n", agenda[i])
			return
		}
	}

}
func deleteContacto(id int) error {
	for i := range agenda {
		if agenda[i].ID == id {
			agenda = append(agenda[:i], agenda[i+1:]...)
			fmt.Println("El contacto se eliminó correctamente\n", agenda[i])
			return nil
		}
	}
	return fmt.Errorf("Contacto no encontrado")

}
func buscarContactoNombre(nombre string) []Contacto {
	var resultados []Contacto
	for _, contacto := range agenda {
		if strings.Contains(strings.ToLower(contacto.Nombre), strings.ToLower(nombre)) {
			resultados = append(resultados, contacto)
		}
	}
	return resultados
}

func buscarContactoNumero(numero int) []Contacto {
	var resultados []Contacto
	for _, contacto := range agenda {
		if contacto.Telefono == numero {
			resultados = append(resultados, contacto)
		}
	}
	return resultados
}
func listContactos() {
	for _, contacto := range agenda {
		fmt.Println("ID |", "NOMBRE |", "TELEFONO |", "CORREO |", "DIRRECION |\n", contacto.ID, "|", contacto.Nombre, "|", contacto.Telefono, "|", contacto.Correo, "|", contacto.Direccion)
	}
	fmt.Println("\n")
}

func main() {
	addContacto("Tomas", 3548515282, "schetot@gmail.com", "felix bogado 1111")
	addContacto("Lionel", 354128479, "lionelgmail.com", "siempreviva 207")
	addContacto("Angel", 3764448877, "angel@yahoo.com", "comandate espora 144")
	listContactos()
	deleteContacto(1)
	listContactos()
	updateContacto(3, "Itachi", 0115545577, "uchiha@yahoo.com", "konoha 123")
	listContactos()

	fmt.Println("resultados de la busqueda por el nombre de Itachi: ")
	resultadosNombre := buscarContactoNombre("itachi")
	for _, resultado := range resultadosNombre {
		fmt.Println("ID |", "NOMBRE |", "TELEFONO |", "CORREO |", "DIRRECION |\n", resultado.ID, "|", resultado.Nombre, "|", resultado.Telefono, "|", resultado.Correo, "|", resultado.Direccion)
	}
	fmt.Println()
	//busqueda del contacto por nombre

	fmt.Println("Busqueda del contacto por numero 354128479: ")
	resultadosNumero := buscarContactoNumero(354128479)
	for _, resultado := range resultadosNumero {
		fmt.Println("ID |", "NOMBRE |", "TELEFONO |", "CORREO |", "DIRRECION |\n", resultado.ID, "|", resultado.Nombre, "|", resultado.Telefono, "|", resultado.Correo, "|", resultado.Direccion)
	}
	fmt.Println()
	//busqueda del contacto por numero

}
