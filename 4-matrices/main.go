package main

import "fmt"

func main() {
	//? Declarar una matríz de 5 elementos int
	var matriz [5]int
	matriz[0] = 10
	matriz[1] = 8
	fmt.Println(matriz)

	//? Declarar una matríz y darle valores a los elementos
	var a = [5]int{0, 2, 4, 6, 8}
	fmt.Println(a)

	//? Declarar una matríz de un número indefinido de elementos
	var b = [...]int{1, 2}
	b[0] = 10
	b[1] = 20

	fmt.Println(b[1])

	//? Recorrer una matríz
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	//? Agregar valores a la matríz
	// b[2] = 30
	// b[3] = 40
	// b[4] = 50

	//? recorrer la matriz usando range
	for index, value := range b {
		fmt.Printf("El índice %d tiene un valor de %d\n", index, value)
	}

	// Matríz bidimensional
	var matrizb = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrizb)

}
