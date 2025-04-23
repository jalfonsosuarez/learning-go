package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewereintheinternet.com",
		"https://graph.microsoft.com",
	}

	//? Crear un canal
	ch := make(chan string)

	for _, api := range apis {
		//! Aplicar la concurrencia con go
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		//? <-ch devuelve datos del canal
		fmt.Println(<-ch)
	}

	//! Espera 2 segundos
	// time.Sleep(5 * time.Second)

	elapsed := time.Since(start)

	fmt.Printf("Tiempo transcurrido: %v", elapsed.Seconds())
}

func checkApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		//? ch <- envía datos al canal
		ch <- fmt.Sprintf("ERROR: ¡%s está caído!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)

}
