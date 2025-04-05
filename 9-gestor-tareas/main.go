package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completada  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

func (l *ListaTareas) marcarCompletada(index int) {
	l.tareas[index].completada = true
}

func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//? Crear una instancia de la lista
	lista := ListaTareas{}

	//? Crear una instancoia de bufio para leer el teclado
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println(
			"Seleccione una opción:\n",
			"1. Nueva tarea.\n",
			"2. Marcar como completada.\n",
			"3. Editar tarea.\n",
			"4. Eliminar tarea.\n",
			"5. Salir.",
		)
		fmt.Print("Introduzca opción:")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Introduzca el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Introduzca la descripción de la tarea:")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("La tarea se añadió correctamente.")
		case 2:
			var index int
			fmt.Print("Introduzca el índice de la tarea:")
			fmt.Scanln(&index)
			lista.marcarCompletada(index)
			fmt.Println("La tarea se marcó como completada correctamente.")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Introduzca el índice de la tarea:")
			fmt.Scanln(&index)
			fmt.Print("Introduzca el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Introduzca la descripción de la tarea:")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("La tarea se mopdificó correctamente.")
			lista.editarTarea(index, t)
		case 4:
			var index int
			fmt.Print("Introduzca el índice de la tarea:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("La tarea se eliminó correctamente.")
		case 5:
			fmt.Println("Saliendo del programa ....")
			return
		default:
			fmt.Println("Opciónn incorrecta.")
		}
		fmt.Println("Lista de tareas")
		printSeparador()
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - Completada: %t\n", i, t.nombre, t.completada)
		}
		printSeparador()
	}
}

func printSeparador() {
	fmt.Println("======================================")
}
