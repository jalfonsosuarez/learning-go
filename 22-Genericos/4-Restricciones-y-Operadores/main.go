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

func main() {

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}

	fmt.Println(Includes(strings, "b"))
	fmt.Println(Includes(strings, "q"))

	fmt.Println(Includes(numbers, 3))
	fmt.Println(Includes(numbers, 9))

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))

	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

}
