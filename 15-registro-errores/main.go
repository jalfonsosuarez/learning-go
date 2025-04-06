package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main(): ")
	log.Print("Mensaje de registro.")
	log.Println("Otro mensaje de registro.")
	//? log.Panic("Aquí se detiene el programa y muestra un seguimiento del error.")
	//! log.Fatal("Este mensaje para el programa.")
	//! fmt.Println("Esto no se ejecutará.")

	//? Crear un fichero para registrar los errores.
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//? Asegura que el archivo quede cerrado al finalizar el programa.
	defer file.Close()

	//? Indicamos el fichero en el que se van a almacenar los errores.
	log.SetPrefix("main(): ")
	log.SetOutput(file)
	log.Print("Registro de errores.")
}
