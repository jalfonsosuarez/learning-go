package main

import "fmt"

func main() {
	// Definiendo un slice (rebanada)
	// var a []ìnt

	// Definir un slice con los días de la semana
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}
	fmt.Println(diasSemana)

	// Obteniendo un slice de diasSemana
	var diasLaborables = diasSemana[1:6]
	fmt.Println(diasLaborables)

	// Obtener el tamaño de un slice
	fmt.Println(len(diasLaborables))
	// Obtener la capacidad delk slice
	fmt.Println(cap(diasLaborables))

	// Añadir elementos al slice
	nuevoSlice := append(diasLaborables, "Sábado")
	fmt.Println(nuevoSlice)

	//? A un slice podemos añadir elementos dinámicamente y también quitar.
	// Elimiar el elemento 2 del slice
	nuevoSlice = append(nuevoSlice[:2], nuevoSlice[3:]...)
	fmt.Println(nuevoSlice)

	//? Creando una rebanada con make
	nombres := make([]string, 5, 10)
	nombres[0] = "Jose"
	nombres[1] = "Juan"
	nombres[2] = "Pedro"
	nombres[3] = "Ramón"
	nombres[4] = "Patricia"

	fmt.Println(nombres)

	//? Copiar elementos de un slice a otro
	nombresCopy := make([]string, 5)
	copy(nombresCopy, nombres)
	fmt.Println(nombresCopy)

}
