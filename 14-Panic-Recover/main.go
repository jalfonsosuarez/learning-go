package main

import "fmt"

func dividir(dividendo, divisor int) {
	//? Se captura el panic con recover y no se detiene la ejecución del programa.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		//? Se genera un panic de forma intencionada.
		panic("No se puede dividir por cero.")
	}
}

func main() {
	dividir(30, 10)
	dividir(50, 5)
	dividir(31, 0)
	dividir(400, 50)
}
