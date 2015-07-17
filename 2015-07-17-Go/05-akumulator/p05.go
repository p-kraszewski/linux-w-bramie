package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type Include struct {
	Z, Do string
}

type Plik struct {
	Baza, Nazwa string
}

var (
	/* RE:
		[Początek linii] 		=   ^
	    [0+ spacji]      		=   \s*
		#                		=   #
		[0+spacji]       		=   \s*
		include           		=   include
		[0+ spacji]      	    =   \s*
		< albo "		 	    =   [<"]
		ZAPAMIĘTYWANY STRING	=   (.*)
		> albo "				=   [>"]
		[0+ dowolnego znaku]    =   .*
		[koniec linii]          =   $
	*/
	re = regexp.MustCompile(`^\s*#\s*include\s*[<"](.*)[>"].*$`)
)

// Na razie każda para jest po prostu wyświetlona na ekranie
func akumulacja(akumulator <-chan Include) {
	for includy := range akumulator {

		fmt.Printf("Plik %s dolacza plik %s\n", includy.Z, includy.Do)
	}
	fmt.Println("Koniec")
}

func analiza(sciezki <-chan Plik, akumulator chan<- Include) {

	for plik := range sciezki {
		wzg, err := filepath.Rel(plik.Baza, plik.Nazwa)
		if err != nil {
			continue
		}

		// Otwarcie pliku READ-ONLY
		plikh, err := os.Open(plik.Nazwa)
		if err != nil {
			continue
		}

		// Opakowanie pliku skanerem
		skaner := bufio.NewScanner(plikh)

		// Przeiterowanie po liniach
		for skaner.Scan() {
			// Sprawdzenie czy linia zawiera include
			includes := re.FindStringSubmatch(skaner.Text())
			if includes != nil {
				// Jak tak, to wysłanie do akumulatora
				akumulator <- Include{wzg, includes[1]}
			}
		}

		// Zamknięcie
		plikh.Close()
	}
	close(akumulator)
}

func skanuj(katalog string, sciezki chan<- Plik) {
	filepath.Walk(katalog,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				sciezki <- Plik{katalog, path}
			}
			return nil
		})

}

func main() {

	// Twarzenie kanału
	var sciezki = make(chan Plik)
	var includy = make(chan Include)
	var baza = "/usr/include"

	// Odpalenie gorutuny (nibywątku) śledzącego kanał
	go akumulacja(includy)
	go analiza(sciezki, includy)

	// Wysłanie ścieżek wszystkich plików do kanału

	skanuj(baza, sciezki)

	// Zamknięcie kanału
	close(sciezki)
	//time.Sleep(3 * time.Second)
}
