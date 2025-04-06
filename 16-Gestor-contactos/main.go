package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// ? Estructura para contactos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// ? Almacena los contactos en un archivo json.
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// ? Carga los contactos desde un archivo json.
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	//? Silice de contacts
	var contacts []Contact

	//? Recuperar los datos desde el fichero
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al recuperar los datos de contactos: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(
			"====== GESTOR DE CONTACTOS =======\n",
			"1. Agrecar contacto.\n",
			"2. Mostrar todos los contactos.\n",
			"3. Salir\n\n",
			"Introduzca una opción: ",
		)
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al introducir la opción: ", err)
			return
		}

		switch option {
		case 1:
			//? Leer todos los datos del contacto.
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			//? Añadimos el contacto al slice de contactos
			contacts = append(contacts, c)

			//? Guardamos en el archivo json.
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar contactos: ", err)
			}
		case 2:
			//? Mostrar todos los contactos.
			fmt.Println("============================================")
			for index, contact := range contacts {
				fmt.Printf("%d.\t Nombre: %s\t Email: %s\t Teléfono: %s\n",
					index, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("============================================")
		case 3:
			return
		default:
			fmt.Println("Opción incorrecta.")
		}
	}

}
