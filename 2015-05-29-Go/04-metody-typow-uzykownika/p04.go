package main

import (
	"fmt"
)

// Własny typ na podstawie int

type Kolor int

// Funkcja IOTA jest równa 0 na początku każdego bloku const i każdej jej użycie
// zwiększa jej wartość o 1 (jest więc generatorem sekwencji 0,1..). IOTA'y można
// użyć w dowolnym kontekście (np 1<<iota dla generowanie potęg dwójki)

// Pierwszy wiersz const MUSI przypisać jakąś wartość i opcjonalnie nadać typ,
// jeżeli jest on inny od wnioskowanego (iota normalnie zwraca int)

// Jeżeli w drugim i następnych wierszach są tylko nazwy, powtarzane jest
// wszystko po nazwie poprzedniego wiersza.

const (
	Czarny   Kolor = iota // Wartość 0
	Czerwony              // Kolor = iota   doda się samo, iota ma wartość 1
	Zielony               // iota ma wartość 2, itd
	Zolty
	Niebieski
	Fioletowy
	Blekitny
	Bialy
)

// Dla każdego typu możemy zdefiniować funcje (metody) na nim działające.
// Metoda "String() string" służy do wyświetlenia typu na ekranie

// func (zm TYP) Nazwa(Parametry) Zwracane {} - praca na kopii
// func (zm *TYP) Nazwa(Parametry) Zwracane {} - praca na oryginale

func (k Kolor) String() string {
	switch k {
	case Czarny:
		return "czarny"
	case Czerwony:
		return "czerwony"
	case Zielony:
		return "zielony"
	case Zolty:
		return "zolty"
	case Niebieski:
		return "niebieski"
	case Fioletowy:
		return "fioletowy"
	case Blekitny:
		return "blekitny"
	case Bialy:
		return "bialy"
	default:
		return "????"
	}
}

// UWAGA! Nie da się zdefiniować metod poza pakietem, w którym definiowany jest
// typ na którym pracują (jak np w Rubym). Czyli do typu string nie da się dodać
// nic swojego

func main() {

	// Automatyczny typ Kolor na podstawie typu stałej
	mojk := Zielony

	// Automatyczne użycie funkcji mojk.String()
	fmt.Println("Kolor=", mojk)
}
