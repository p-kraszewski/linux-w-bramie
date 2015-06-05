package main

import (
	"fmt"
)

var (
	t1 = [5]byte{}                // Statyczna tablica 5 bajtów o wartości 0
	t2 = [...]byte{1, 2, 3, 4, 5} // Statyczna tablica 5 bajtów

	t3 = []byte{}              // Dynamiczna tablica
	t4 = []byte{1, 2, 3, 4, 5} // Dynamiczna tablica 3 bajtów (co innego niż t2!)

	t5 = t2[2:4] // Wycinek tablicy t2 zawierający wyrazy od 2 do 4 (bez 4)
)

func main() {
	fmt.Println("t1", t1)
	fmt.Println("t2", t2)
	fmt.Println("t3", t3)
	fmt.Println("t4", t4)

	fmt.Println("t5", t5)

	// Modyfikacja slice'a modyfikuje oryginalną tablicę

	t5[0] = 100
	fmt.Println("t5'", t5)
	fmt.Println("t2'", t2)

	// Modyfikacja tablicy widoczna jest w slice'u
	t2[3] = 200
	fmt.Println("t5''", t5)
	fmt.Println("t2''", t2)

}
