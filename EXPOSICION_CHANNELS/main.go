package main

import (
	"fmt"
	"time"
)

// Función que envía números a un unbuffered channel
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Enviado: %d\n", i)
		time.Sleep(time.Second)
	}
	close(ch) // Buena práctica: cerrar el channel cuando no se usarán más datos
}

// Función que envía mensajes a un buffered channel
func sendMessages(ch chan string) {
	messages := []string{"Hola", "Mundo", "desde", "Go!"}
	for _, msg := range messages {
		ch <- msg
		fmt.Printf("Mensaje enviado: %s\n", msg)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func main() {
	// Crear un unbuffered channel
	intChan := make(chan int)

	// Crear un buffered channel
	stringChan := make(chan string, 2)

	// Lanzar goroutines
	go sendNumbers(intChan)
	go sendMessages(stringChan)

	// Usar select para manejar múltiples channels
	for {
		select {
		case num, ok := <-intChan:
			if ok {
				fmt.Printf("Número recibido: %d\n", num)
			} else {
				intChan = nil // El channel está cerrado
			}
		case msg, ok := <-stringChan:
			if ok {
				fmt.Printf("Mensaje recibido: %s\n", msg)
			} else {
				stringChan = nil // El channel está cerrado
			}
		}

		// Salir del bucle cuando ambos channels estén cerrados
		if intChan == nil && stringChan == nil {
			break
		}
	}

	fmt.Println("Todos los datos han sido procesados.")
}
