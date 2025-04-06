package main

import "fmt"

func main() {
	//? La función precedida de defer, se ejecutará al final.
	defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)

	//? Cuando hay varias funciones precedidas de defer, la ejecución se añade a una pila
	//? de tal forma que la primera en entrar a la pila será la última en ejecutarse.
	//? La primera en llegar es la que imprime el número 3.
	defer fmt.Println(6)
	defer fmt.Println(5)
	defer fmt.Println(4)
}
