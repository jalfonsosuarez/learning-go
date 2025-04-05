package main

import (
	"fmt"

	"rsc.io/quote"
)

// Contantes
// Hay que declararlas e inicializar su valor.
// Declarada fuera de una función es visible por todo el paquete.
const PI float32 = 3.141592

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	// Declaración de variables
	// var firstName, lastName string
	// var age int

	// Otra forma de declarar variables
	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )
	// firstName = "Jose Alfonso"
	// lastName = "Suárez Moreno"
	// age = 60

	// Otra forma de declarar variables y asignarle valores
	// Las variables definidas dentro de una función sólo son visibles por la función,
	var (
		firstName string = "Jose Alfonso"
		lastName  string = "Suárez Moreno"
		age       int    = 60
	)

	fmt.Println(firstName, lastName, age)

}
