package main

import "fmt"

// ? Una función variádica recibe un número inteterminado de parámetros
// ? Espera un número indeterminado de valores tipo int
func suma(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// ? Espera un número indeterminado de valores de cualquier tipo
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	fmt.Println(suma(2, 3, 4, 55, 6, 7, 92, 3, 4))
	imprimirDatos("Cadena", true, 25, 2.74)
}
