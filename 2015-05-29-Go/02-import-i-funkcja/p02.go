package main

/*
	Biblioteki importujemy w bloku "import". Kod źródłowy biblioteki BIB szukany
	jest w katalogu $GOROOT/src/BIB, potem $GOPATH/src/BIB

	Skompilowane biblioteki są w plikach odpowiednio $GOROOT/pkg/ARCH/BIB.a i
	$GOPATH/pkg/ARCH/BIB.a, gdzie ARCH to kombinacja OS i architektury na jaką
	skompilowano bibliotekę (tak, Go natywnie wspiera kross-kompilację między
	systemami i architekturami, na Malince da się skompilować 64-bitowy EXE dla
	Windows), na przykład linux_amd64

	Biblioteki mogą być hierarchiczne (np "net/http"), jednak nie implikuje to
	żadnych zależności - "net/http" nie ma żadnego szczególnego dostępu do "net"
	i odwrotnie.
*/

import (
	"fmt"
)

/*
 Typy danych podobne do C/C++:
 int/uint, int8/uint8..16..32..64
 float32/64, complex64/128
 bool
 string

 i parę egzotyków

*/

// funkcje, metody, zmienne, stałe, cokolwiek zaczynającego się od wielkiej
// litery są publiczne, od małej są prywatne (dla pakietu)

// Wszystkie deklaracje są w kolejności nazwa-typ nie typ-nazwa (jak w C)

func Silnia(n uint64) uint64 {

	// if, for itd nie mają nawiasów, jak w C. Klamra jest obowiązkowa i musi
	// być w linii z operatorem

	if n <= 1 {
		return 1
	}

	return n * Silnia(n-1)
}

func main() {

	// Wywołanie funkcji z biblioteki - przez kropkę. Nawet jeżeli biblioteka
	// jest zagnieżdżona "moj_projekt/biblioteka/narzedzia", przed kropką jest
	// tylko ostatnie część (czyli np funkcja narzedzia.testuj() )
	//
	// Funkcje z bieżącego modułu można wołać bez kwalifikacji (samo Silnia() )
	fmt.Println("20!=", Silnia(20))
}
