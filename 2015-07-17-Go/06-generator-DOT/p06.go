package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type Include struct {
	Z, Do string
}

type Plik struct {
	Baza, Nazwa string
}

type Nic struct{}
type ZbiorStr map[string]Nic

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
	re = regexp.MustCompile(`^\s*#\s*include\s*([<"])([^>"]*).*$`)

	kolor = []string{"", "red", "green", "black"} // To NIESTETY nie jest CONST

	czekaj sync.WaitGroup
)

// Zbieranie danych i generacja pliku wynikowego
func akumulacja(akumulator <-chan Include) {

	// Zgłaszamy że jesteśmy niegotowi
	czekaj.Add(1)

	// Funkcja uruchomiona przy opuszczeniu zakresu dowolną drogą
	// i wtedy zgłaszamy, że jesteśmy gotowi
	defer czekaj.Done()

	// Hash węzłów z ich typem (1=wchodzący, 2=wychodzący, 3=oba)
	wezly := make(map[string]byte)

	// Hash połączeń - typ hash(wezel źródłowy->zbiór węzłów docelowych)
	linki := make(map[string]ZbiorStr)

	// Że jeszcze żyję
	fmt.Println("Zbieram dane")

	// Każdy nowy link
	for includy := range akumulator {

		// Stwórz lub uaktualnij węzeł docelowy
		wezly[includy.Do] |= 1

		// Stwórz lub uaktualnij węzeł źródłowy
		wezly[includy.Z] |= 2

		// Jeżeli jeszcze togo węzła źródlowego nie ma w bazie linków, to go
		// stwórz
		if _, jest := linki[includy.Z]; !jest {
			linki[includy.Z] = make(ZbiorStr)
		}

		// I dopisz nowy link
		linki[includy.Z][includy.Do] = Nic{}
	}
	fmt.Println("Generuje")

	// Otwarcie pliku do zapisu
	file, err := os.Create("graf.dot")
	if err != nil {
		return
	}

	// Nagłówek pliku DOT
	fmt.Fprintln(file, "digraph includy{")
	fmt.Fprintln(file, "splines=ortho;")

	// Wszystkie węzły mają kształt prostokąta i odpowiedni kolor ramki
	for plik, typ := range wezly {
		// Kolory zależą od typu węzła

		fmt.Fprintf(file, "\"%s\" [shape=box color=%s];\n", plik, kolor[typ])
	}

	// Pusta linia
	fmt.Fprintln(file, "")

	// Dorobienie strzałek
	for zrodlo, cele := range linki {
		for cel, _ := range cele {
			fmt.Fprintf(file, "\"%s\" -> \"%s\";\n", cel, zrodlo)
		}
	}

	// Koniec pliku DOT
	fmt.Fprintln(file, "}")

	// Zamknięcie
	file.Close()

}

func analiza(sciezki <-chan Plik, akumulator chan<- Include) {

	for plik := range sciezki {
		// Normalizacja ścieżek względam ".usr/include" a nie względem
		wzg, err := filepath.Rel(`/usr/include`, plik.Nazwa)
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
				z := wzg
				do := includes[2]

				// Normalizujemy ścieżki względne
				if strings.HasPrefix(do, `.`) || includes[1] == `"` {
					cwd := filepath.Dir(z)
					do = filepath.Clean(cwd + `/` + do)
				}

				// Wysyłąmy do akumulatora
				akumulator <- Include{z, do}
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
	var baza = "/usr/include/openssl"

	// Odpalenie gorutuny (nibywątku) śledzącego kanał
	go akumulacja(includy)
	go analiza(sciezki, includy)

	// Wysłanie ścieżek wszystkich plików do kanału

	skanuj(baza, sciezki)

	// Zamknięcie kanału
	close(sciezki)
	czekaj.Wait()
}
