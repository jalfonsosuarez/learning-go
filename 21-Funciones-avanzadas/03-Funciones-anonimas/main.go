package main

import "fmt"

// ? Esta función recibe otra como parámetro
func saludar(nombre string, f func(string)) {
	f(nombre)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {
	//? Función anónima
	func() {
		fmt.Println("Hola desde una función anónima.")
	}()

	//? Función anónima asignada a una variable
	saludo := func(nombre string) {
		fmt.Printf("Hola %s\n", nombre)
	}

	//? La variable se comporta comu una función
	saludo("Pepe")

	//? Se invoca a la función anónima asignada a la valiable saludo
	saludar("Pedro", saludo)

	fmt.Println(aplicar(duplicar, 4))
	fmt.Println(aplicar(triplicar, 9))
}
