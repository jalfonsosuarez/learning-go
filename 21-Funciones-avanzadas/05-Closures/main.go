package main

import "fmt"

// ? Un closure (o clausura) es una función anónima dentro de otra función
// ? que permite recordar los valores de la función padre
func incremeter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := incremeter()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
