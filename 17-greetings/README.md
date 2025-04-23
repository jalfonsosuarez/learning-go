# Saludos en Go

Paquete que proporciona saludos personalizados.

Curso Udemy:

[GOLANG: Curso profesional de Go - De cero a Master 2025](https://www.udemy.com/course/curso-golang/)

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/jalfonsosuarez/greetings
```

## Uso

Ejemplo de como usar el paquete en tu proyecto Go:

```go
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
```
