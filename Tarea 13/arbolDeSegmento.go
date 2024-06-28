/*Implementa un árbol de segmentos en Go que admita dos operaciones: actualizar un valor en un índice específico del arreglo subyacente
y calcular la suma de un rango de valores en el arreglo. Proporciona ejemplos de uso para demostrar la eficiencia de tu implementación.*/

package main

import (
	"fmt"
	"math"
)

type ArbolSegmento struct {
	n       int
	arbol   []int
	entrada []int
}

func NuevoArbolSegmento(arr []int) *ArbolSegmento {
	n := len(arr)
	tamArbol := 2*int(math.Pow(2, math.Ceil(math.Log2(float64(n))))) - 1
	arbol := make([]int, tamArbol)
	as := &ArbolSegmento{n: n, arbol: arbol, entrada: make([]int, n)}
	copy(as.entrada, arr)
	as.construir(0, 0, n-1)
	return as
}

func (as *ArbolSegmento) construir(nodo, inicio, fin int) {
	if inicio == fin {
		as.arbol[nodo] = as.entrada[inicio]
	} else {
		medio := (inicio + fin) / 2
		hijoIzq := 2*nodo + 1
		hijoDer := 2*nodo + 2
		as.construir(hijoIzq, inicio, medio)
		as.construir(hijoDer, medio+1, fin)
		as.arbol[nodo] = max(as.arbol[hijoIzq], as.arbol[hijoDer])
	}
}
func (as *ArbolSegmento) actualizar(idx, nuevoVal int) {
	diferencia := nuevoVal - as.entrada[idx]
	as.entrada[idx] = nuevoVal
	as.actualizarAux(0, 0, as.n-1, idx, diferencia)
}

func (as *ArbolSegmento) actualizarAux(nodo, inicio, fin, idx, diferencia int) {
	if idx < inicio || idx > fin {
		return
	}
	as.arbol[nodo] += diferencia
	if inicio != fin {
		medio := (inicio + fin) / 2
		hijoIzq := 2*nodo + 1
		hijoDer := 2*nodo + 2
		as.actualizarAux(hijoIzq, inicio, medio, idx, diferencia)
		as.actualizarAux(hijoDer, medio+1, fin, idx, diferencia)
	}
}

func (as *ArbolSegmento) maximo(l, r int) int {
	return as.maximoAux(0, 0, as.n-1, l, r)
}

func (as *ArbolSegmento) maximoAux(nodo, inicio, fin, l, r int) int {
	if r < inicio || l > fin {
		return math.MinInt64
	}
	if l <= inicio && r >= fin {
		return as.arbol[nodo]
	}
	medio := (inicio + fin) / 2
	hijoIzq := 2*nodo + 1
	hijoDer := 2*nodo + 2
	maxIzq := as.maximoAux(hijoIzq, inicio, medio, l, r)
	maxDer := as.maximoAux(hijoDer, medio+1, fin, l, r)
	return max(maxIzq, maxDer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 3, 2, 7, 9, 11}
	as := NuevoArbolSegmento(arr)

	fmt.Println("Maximo en el rango [1, 3]:", as.maximo(1, 3))
	fmt.Println("Máximo en el rango [2, 5]:", as.maximo(2, 5))

	as.actualizar(3, 4)
	fmt.Println("Maximo en el rango [1, 3] después de actualizar el índice 3 a 4:", as.maximo(1, 3))
}
