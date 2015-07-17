package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	/* RE:
		[Początek linii] 		=   ^
	    [0+ spacji]      		=   \s*
		#                		=   #
		[0+spacji]       		=   \s*
		include           		=   include
		[0+ spacji]      	    =   \s*
		< albo "		 	    =   [<"]
		ZAPAMIĘTYWANY nie ">	=   ([^>"]*)
		[0+ dowolnego znaku]    =   .*
		[koniec linii]          =   $
	*/
	re = regexp.MustCompile(`^\s*#\s*include\s*[<"]([^>"]*).*$`)
)

func skanuj(baza, plik string) {
	wzg, err := filepath.Rel(baza, plik)
	if err != nil {
		return
	}

	fmt.Printf("Skanuję plik %s.\n", wzg)

	// Otwarcie pliku READ-ONLY
	plikh, err := os.Open(plik)
	if err != nil {
		return
	}

	// Opakowanie pliku skanerem
	skaner := bufio.NewScanner(plikh)

	// Przeiterowanie po liniach
	for skaner.Scan() {
		// Sprawdzenie czy linia zawiera include
		includes := re.FindStringSubmatch(skaner.Text())
		if includes != nil {
			// Jak tak, to wypisanie go
			fmt.Println(includes[1])
		}
	}

	// Zamknięcie
	plikh.Close()
}

func main() {
	skanuj("/usr/include/", "/usr/include/stdio.h")
}
