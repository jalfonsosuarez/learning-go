package main

import (
	"fmt"
	"os"
)

func main() {
	//? Crear el fichero hola.txt
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	//? Con defer, se cerrará el fichero al terminar la ejecución del programa.
	defer file.Close()

	_, err = file.Write([]byte("Hola José"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
