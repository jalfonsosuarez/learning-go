package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// go 1.19 o anterior rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(100))
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		fmt.Printf("Ingrese un número entre 0 y 100 (intentos restantes %d): ", maxIntentos-intentos)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Printf("Has acertado, el número oculto era %d\n", numAleatorio)
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("Pruebe con un número mayor")
		} else {
			fmt.Println(("Pruebe un número menor"))
		}
		intentos++
	}

	fmt.Printf("Ya no hay más intentos. El número oculto era %d\n", numAleatorio)
	jugarNuevamente()
	return
}

func jugarNuevamente() {
	var respuesta string
	fmt.Print("¿Quieres volver a jugar (s/n): ")
	fmt.Scanln(&respuesta)

	switch respuesta {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar.")
	default:
		fmt.Println("Elección no válida, prueba de nuevo.")
		jugarNuevamente()
	}
}
