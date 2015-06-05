// Komentarze jak w C/C++ (czyli // i /* */)

/*
 Każdy plik Go musi należeć do jakiegoś pakietu.
 Pliki wykonywalne (aplikacje) muszą należeć do pakietu "main"

 W jednym katalogu może znajdować się tylko jeden plik zawierający funkcję main.

 Generalnie 1 katalog == 1 biblioteka albo 1 aplikacja
*/
package main

/*
 Funkcje deklaruje się przez słowo kluczowe "func", po którym następuje nazwa
 funkcji, parametry funkcji (tutaj puste) i wartość/ci zwracane (tutaj też puste)

 Klamra MUSI znajdować się w linii, której dotyczy - niezależnie od tego, czy
 jest to func, for, if czy cokolwiek innego z klamrami
*/

func main() {
	// Funkcja nic nie robi, nic nie zwraca
}
