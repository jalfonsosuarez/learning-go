package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	//? Crear la estructura de forma habitual
	//? Para poder crearlo así los atributos deben ser públicos (primera letra mayúscula)
	// var myBook = book.Book{
	// 	Title:  "Cien mil leguas de viaje submarino",
	// 	Author: "Julio Verne",
	// 	Pages:  300,
	// }

	// myBook.PrintInfo()

	//? Crear la estructura usando el construtor
	myBook := book.NewBook(
		"Viaje al centro de la tierra",
		"Julio Verne",
		300,
	)

	myBook.PrintInfo()

	//? Modificando el valor de un atributo
	myBook.SetTile("Cien mil leguas de viaje submarino")
	myBook.PrintInfo()

	//? Accediendo al valor de un atributo
	fmt.Println("Title: ", myBook.GetTitle())

	//? Crear un objeto de tipo TextBook
	myTextBook := book.NewTextBook("Matemáticas", "Eduardo López", 223, "Everest", "1º EP")
	myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)

	myPerro := animal.Perro{Nombre: "Chaman"}
	myGato := animal.Gato{Nombre: "Sirena"}

	animal.HacerSonido(&myPerro)
	animal.HacerSonido(&myGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Chaman"},
		&animal.Gato{Nombre: "Sirena"},
		&animal.Perro{Nombre: "Coco"},
		&animal.Gato{Nombre: "Silvestre"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
