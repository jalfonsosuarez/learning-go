package main

import (
	"fmt"
)

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// return 0, errors.New("No se puede dividir por cero.")
		return 0, fmt.Errorf("No se puede dividir por cero.")
	}
	return dividendo / divisor, nil
}

func main() {
	result, error := dividir(10, 0)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Resultado: ", result)
	}

	result, error = dividir(10, 2)
	if error != nil {
		fmt.Println("Error: ", error)
		return
	} else {
		fmt.Println("Resultado: ", result)
	}
}
