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

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

type Product[T uint | string] struct {
	Id         T
	Desciption string
	Precio     float32
}

func main() {

	product1 := Product[uint]{1, "Zapato", 35}
	product2 := Product[string]{"adf43-456-asfa665", "Camisa", 65}

	fmt.Println(product1)
	fmt.Println(product2)

}
