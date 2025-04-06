package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Numero:", num)

	//? Forzamos el error
	str = "123A"
	num, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Numero:", num)
}
