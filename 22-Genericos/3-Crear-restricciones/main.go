package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// ? Se define el tipo integer que deriva de int.
type integer int

// ? Se define una interface para las restrincciones
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// ? Se usan los tipos definidos en constrainst.
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	//? Se indica el tipo de datos de los parámetros que se le pasan a la función.
	fmt.Println(Sum[uint](4, 5, 6, 9, 4, 15))

	fmt.Println(Sum(4.5, 6, 9.4, 15))
	var num1 integer = 100
	var num2 integer = 200
	fmt.Println(Sum(num1, num2))
}
