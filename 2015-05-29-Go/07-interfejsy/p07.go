package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
 Go ma interfejsy. Deklaracja wygląda na przykład tak (pakiet "io"):

type Reader interface {
        Read(p []byte) (n int, err error)
}

Jeżeli jakikolwiek (!) typ ma zdefiniowaną metodę o powyższej sygnaturze, to
można go użyć wszędzie tam, gdzie oczekiwany jest interfejs io.Reader .

*/

// Nasz wymyślony typ
type Wypelniaczka byte

// Dla tego typu definiujemy metodę o sygnaturze jak wyżej. Zachowuje się jak
// nieskończone źródło bajtów b

func (b *Wypelniaczka) Read(p []byte) (n int, err error) {

	// pos będzie miał wartości indeksów kolejnych elementów p
	for pos := range p {
		// Nadpisujemy p[pos] - musimy zrobić rzutowanie z Wypelniaczka na byte
		p[pos] = byte(*b)
	}

	// Zwracamy długość "odczytanych" danych i nil jako brak błędu
	return len(p), nil
}

func main() {
	var b Wypelniaczka = 42 //  Bo tak

	// slice bajtów
	bufor := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Przed odczytem=", bufor)

	b.Read(bufor) // bufor jest blokiem, do którego odczytuje się dane

	fmt.Println("Po odczycie=", bufor)

	// Zrobimy sobie zatem bufor wyjściowy implementujący io.Writer:

	var do_zapisu bytes.Buffer

	// I skopiujmy do niego 16 bajtów z b. Pobieranie jest z _interfejsu_ Reader.
	// Nasz typ ten interfejs spełnia, więc można go użyć w funkcji CopyN

	// func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

	io.CopyN(&do_zapisu, &b, 16)

	fmt.Println("Skopiowany bufor=", do_zapisu.Bytes())
}
