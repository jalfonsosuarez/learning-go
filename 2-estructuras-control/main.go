package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Es por la mañana")
	} else if hora < 17 {
		fmt.Println("Es por la tarde")
	} else {
		fmt.Println("Es por la noche")
	}

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> Mac OS X")
	default:
		fmt.Println("Go run -> os")
	}

	var i int

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		if j == 5 {
			break
		}
	}

	for j := 0; j < 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}

	funcion()

	hello("Jose")

	fmt.Println("Hola", otraFunc("Patricia"))

	fmt.Println(calc(2, 4))

}

func funcion() {
	fmt.Println("Hola, soy una función en GO!")
}

func hello(name string) {
	fmt.Printf("Hola %s desde la función hello!\n", name)
}

func otraFunc(name string) string {
	return name
}

func calc(a, b int) (int, int) {
	sum := a + b
	mult := a * b
	return sum, mult
}
