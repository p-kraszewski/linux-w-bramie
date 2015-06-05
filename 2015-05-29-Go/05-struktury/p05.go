package main

import (
	"fmt"
)

// Typ Punkt to punkt w 3 wymiarach.

type Punkt struct {
	// Z małej litery, więc gdyby całość była w bibliotece, użytkownik mógłby
	// robić zmienne typu Punkt, ale nie mógłby zajrzeć do środka (składowe
	// prywatne poziomu pakietu)
	x, y, z float64
}

// Wyświetlajka (pracuje na kopii)
func (p Punkt) String() string {
	// Mamy dostęp do prywatnych składowych x, y i z, bo jesteśmy w tym samym
	// pakiecie

	return fmt.Sprintf(
		"<%+3.2f,%+3.2f,%+3.2f>",
		p.x, // Dostęp do składowych jak w C
		p.y,
		p.z, // <--- Przecinek na końcu!!!
	)
}

func main() {
	// Tworzymy zmienną typu punkt z domyślnymi wartościami
	p1 := Punkt{}

	// Tworzymy zmienną typu punkt z wskazanymi wartościami
	p2 := Punkt{1, 2, 3}

	// Tworzymy zmienną typu punkt z konkretną jedną wartością - reszta domyślne
	p3 := Punkt{z: 5}

	// Tworzymy zmienną typu *punkt
	p4 := &Punkt{z: 7, x: 1}

	var p5 Punkt
	var p6 *Punkt // Pusty wskaźnik!!!!

	fmt.Println("Punkt1=", p1)
	fmt.Println("Punkt2=", p2)
	fmt.Println("Punkt3=", p3)
	fmt.Println("Punkt4=", p4) // Automatyczna dereferencja
	fmt.Println("Punkt5=", p5)
	fmt.Println("Punkt6=", p6) // Pokazuje <nil>

	p6 = new(Punkt)

	fmt.Println("Punkt6'=", p6) // Pokazuje <0,0,0>

}
