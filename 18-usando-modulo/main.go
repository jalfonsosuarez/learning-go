package main

import (
	"fmt"
	"log"

	"github.com/jalfonsosuarez/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	//! Elimina la fecha y la hora del log de errores.
	log.SetFlags(0)

	message, err := greetings.Hello("Jose")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{
		"Jose",
		"Patty",
		"Juan",
		"Ana",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
