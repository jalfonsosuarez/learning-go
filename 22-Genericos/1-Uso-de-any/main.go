package main

import "fmt"

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	PrintList("Jose", 1, 2.5, true)
}
