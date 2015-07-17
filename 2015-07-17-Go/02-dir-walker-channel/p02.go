package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Funkcja czeka na kolejne dane z kanału i wyświetla je na ekranie. Po
// zamknięciu kanału wyświetla napis "Koniec" i kończy swoje działanie
func analiza(pliki <-chan string) {
	for plik := range pliki {
		fmt.Printf("Plik: %s\n", plik)
	}
	fmt.Println("Koniec")
}

func main() {

	// Twarzenie kanału
	var sciezki = make(chan string)

	// Odpalenie gorutuny (nibywątku) śledzącego kanał
	go analiza(sciezki)

	// Wysłanie ścieżek wszystkich plików do kanału
	filepath.Walk("/usr/include",
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				sciezki <- path
			}
			return nil
		})

	// Zamknięcie kanału
	close(sciezki)
}
