package main

import "fmt"

// ? Se define el tipo integer que deriva de int.
type integer int

// ? Especificamos los tipos con los que va a trabajar la función.
// ? El símbolo ~ sirve para indicar que se usen los del tipo padre y sus derivados.
func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(Sum(4.5, 6, 9.4, 15))
	var num1 integer = 100
	var num2 integer = 200
	fmt.Println(Sum(num1, num2))
}
