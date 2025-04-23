package main

import (
	"fmt"
)

// ? Definir uns estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

// ? Declaración de un método para la estructuta Persona
// ? Indicamos en p un puntero que apunta a la direción de memoria donde está Persona
func (p *Persona) saluda() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	//? Declaramos una variable p que será un puntero a la dirección de memoria de x
	//? x y p tendrán acceso a la misma dirección de memoria
	var p *int = &x

	//? Imprime el valor de x
	fmt.Println(x)
	//? Imprime la dirección de memoria de x
	fmt.Println(&x)
	//? Imprime la direccion de memoria de x almacenado en p
	fmt.Println(p)

	//? Enviar la direccion de memoria a la función
	editar(&x)
	fmt.Println(x)

	//? Crear una instancoa de Persona
	p1 := Persona{
		nombre: "José",
		edad:   60,
		correo: "jose@mail.com",
	}
	p1.saluda()
}

// ? La función recibe un puntero
func editar(x *int) {
	//? Modifica el valor de la dirección
	*x = 20
}
