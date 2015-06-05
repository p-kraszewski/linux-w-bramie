package main

import (
	"fmt"
	"sync"
	"time"
)

// Go jest współbieżny

var wg sync.WaitGroup

func NieCalkiemWatek(n int) {
	fmt.Println("Poczatek gorutyny", n)
	time.Sleep(time.Second)
	fmt.Println("Koniec gorutyny", n)

	// Kasujemy oczekiwanie
	wg.Done()
}

func main() {
	// Synchronizacja wątków

	przedtem := time.Now()

	// 4 "wątki" o czasie wykonania 1 sekunda
	for i := 1; i <= 4; i++ {
		// Dodajemy oczekiwanie
		wg.Add(1)
		// Odpalenie funkcji w groutynie. To może ale nie musi być wątek.
		go NieCalkiemWatek(i)
	}

	// Czekamy aż oczekiwanie dojdzie do 0 (w sumie semafor...)
	wg.Wait()

	potem := time.Now()

	fmt.Println("Czas wykonania 4 gorutyn", potem.Sub(przedtem))

}
