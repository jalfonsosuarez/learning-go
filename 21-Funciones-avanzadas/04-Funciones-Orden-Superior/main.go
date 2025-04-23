package main

import "fmt"

// ? Una función de orden superior es la que recibe una función como parámetro y retorna otra función
// ? Con esta característica se pueden definir funciones personalizadas en tiempo de ejecución
func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

func main() {
	r := double(addOne, 3)
	fmt.Println(r)
}
