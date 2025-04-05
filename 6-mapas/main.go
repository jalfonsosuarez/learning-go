package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	//? Muestra todos los elementos del mapa.
	fmt.Println(colors)
	//? Muestra el valor para el elemento "rojo",
	fmt.Println(colors["rojo"])
	//? Añadir un nuevo elemento al mapa.
	colors["negro"] = "#000000"
	//? Asignar el valor de un elemento a una variable.
	miVariable := colors["azul"]
	fmt.Println(miVariable)
	//? Comprobar si existe el elemento en el mapa
	otraVariable, ok := colors["blanco"]
	if ok {
		fmt.Println(otraVariable)
	} else {
		fmt.Println("No existe el elemento blanco en colors.")
	}
	//? Segunda forma de comprobar si existe el elemento en el mapa
	if otraVariable, ok := colors["amarillo"]; ok {
		fmt.Println(otraVariable)
	} else {
		fmt.Println("No existe el elemento amarillo en colors.")
	}
	//? Eliminar un elemento del mapa
	delete(colors, "negro")
	fmt.Println(colors)
	//? Iterenda un mapa
	for clave, valor := range colors {
		fmt.Printf("Clave %s, valor %s\n", clave, valor)
	}
}
