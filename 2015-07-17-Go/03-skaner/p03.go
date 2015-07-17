package main

import (
	"bufio"
	"fmt"
	"os"
)

func skanuj(plik string) {
	fmt.Printf("Skanuję plik %s.\n", plik)

	// Otwarcie pliku READ-ONLY
	plikh, err := os.Open(plik)
	if err != nil {
		return
	}

	// Opakowanie pliku skanerem
	skaner := bufio.NewScanner(plikh)

	// Przeiterowanie po liniach
	for skaner.Scan() {
		// Wypisanie linii
		fmt.Println(skaner.Text())
	}

	// Zamknięcie
	plikh.Close()
}

func main() {
	skanuj("/usr/include/stdio.h")
}
