package main

import (
	"fmt"
	"time"
)

// Producent, konsument, kilka procesorów

var (

	// Kanały transmisji obiektów typu INT

	do_obrobki = make(chan int)
	obrobione  = make(chan int)
)

// Konsument czeka na kanale obrobione i pobiera kolejne obiekty

func Konsument() {
	skonczone := 0
	for obiekt := range obrobione {
		skonczone++
		fmt.Println("Konsument skonsumował obiekt", obiekt, "(", skonczone, "dotąd)")
	}

	fmt.Println("Konsument skonczył pracę")
}

// Procesor czeka na kanale do_obrobki i i przekazuje obiekty do kanału obrobione

func Procesor(nr int) {
	skonczone := 0
	for obiekt := range do_obrobki {
		skonczone++
		fmt.Println("Procesor", nr, "pobrał obiekt", obiekt)
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Procesor", nr, "obrobił obiekt", obiekt)
		obrobione <- obiekt
	}
	fmt.Println("Producent", nr, "skonczył pracę, przetworzyl", skonczone)

}

func main() {
	// Odpalenie konsumenta
	go Konsument()

	// Odpalenie producentów
	for proc := 1; proc <= 5; proc++ {
		go Procesor(proc)
	}

	// Wyprodukowanie i "wypchnięcie" obiektów
	for obj := 1; obj <= 25; obj++ {
		fmt.Println("Producent dostarcza obiekt", obj)
		do_obrobki <- obj
		fmt.Println("Producent dostarczył obiekt", obj)
	}

	// Czekamy (brak synchronizacji
	time.Sleep(time.Second * 2)
	close(do_obrobki)
	close(obrobione)

	// Czekamy (brak synchronizacji
	time.Sleep(time.Second * 1)
}
