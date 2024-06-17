/*
Implementar una estructura de datos para almacenar la información de cada producto (nombre, precio y cantidad disponible).
Permitir la adición de nuevos productos con sus respectivas cantidades.
Permitir la actualización de la cantidad disponible de un producto existente.
Permitir la eliminación de productos del inventario.
Mostrar el inventario completo, incluyendo el nombre, precio y cantidad disponible de cada producto.
*/
package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

var inventario []Producto

func addProduct(nombre string, precio float64, cantidad int) {
	id := len(inventario) + 1
	producto := Producto{ID: id, Nombre: nombre, Precio: precio, Cantidad: cantidad}
	inventario = append(inventario, producto)
	fmt.Println("El producto se agregó correctamente al inventario\n", producto)
}
func updateProduct(id int, nombre string, precio float64, cantidad int) error {
	for i := range inventario {
		if inventario[i].ID == id {
			inventario[i].Nombre = nombre
			inventario[i].Precio = precio
			inventario[i].Cantidad = cantidad
			fmt.Println("El producto se actualizó correctamente\n", inventario[i])
			return nil
		}
	}
	return fmt.Errorf("No se encontró el producto")
}

func deleteProduct(id int) error {
	for i, prod := range inventario {
		if prod.ID == id {
			inventario = append(inventario[:i], inventario[i+1:]...)
			fmt.Println("El producto se eliminó correctamente\n", inventario[i])
			return nil
		}
	}
	return fmt.Errorf("No se encontró el producto")
}
func listProduct() {
	fmt.Println("Inventario de productos:\n")
	for _, prod := range inventario {
		fmt.Println("ID | Name | Price | Amount \n", prod.ID, prod.Nombre, prod.Precio, prod.Cantidad)
	}
}
func main() {
	addProduct("Chicle", 100.00, 300)
	addProduct("Chupetin", 90.00, 150)
	addProduct("Alfajor", 200.00, 50)
	listProduct()
	updateProduct(3, "Chocolate", 250.00, 20)
	listProduct()
	deleteProduct(2)
	listProduct()

}
