package main

import (
	"fmt"
)

// Tworzymy własny typ - na podstawie INT

type Kolor int

func main() {
	// Nowa zmienna typu Kolor
	var nowy_kolor Kolor

	// Przypisujemu mu wartość - stała/literał 5 się sam zrzutuje "w górę" na
	// typ Kolor
	nowy_kolor = 5

	// Zapis skrócony:
	//   ZM := WART
	// jest równoważne
	//  var ZM typ_WART
	//  ZM = WART
	//
	// Ale uwaga! inny_kolor jest typu int ("natywnego" dla 6)
	inny_kolor := 6

	// Go jest językiem silnie typowanym, więc nie można zrobić przypisania
	//
	// nowy_kolor = inny_kolor
	//
	// Błąd: cannot use inny_kolor (type int) as type Kolor in assignment
	//
	// Trzeba zrobić jawne rzutowanie (nawias jest odwrotnie niż w C)
	nowy_kolor = Kolor(inny_kolor)

	fmt.Println("Kolor=", nowy_kolor)

}
