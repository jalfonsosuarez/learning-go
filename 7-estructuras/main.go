package main

import "fmt"

// ? Definir uns estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {

	//? Declaramos una variable de tipo Persona y damos valores a sus campos.
	var p1 Persona
	p1.nombre = "José"
	p1.edad = 60
	p1.correo = "joseasuarez@gmail.com"
	fmt.Println(p1)

	//? Otra forma de usar una estructura
	p2 := Persona{
		nombre: "Juan",
		edad:   40,
		correo: "juan@mail.com",
	}
	fmt.Println(p2)

	//? También se pueden asignar valores por posición
	p3 := Persona{
		"Patty",
		54,
		"patty@mail.com",
	}
	fmt.Println(p3)

	//? Acceder a los valores por el nombre del campo
	fmt.Println(p3.nombre)

}
