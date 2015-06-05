package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Funkcja uruchamiana dla obsługi kolejnego zapytania

func handler(w http.ResponseWriter, r *http.Request) {

	// r => handler => w

	fmt.Fprintf(w, "Hi there, I love %s!\n\n", r.URL.Path[1:])

	// Konwersja struktury Request na zapis JSON (biblioteka standardowa Go!)
	ans, _ := json.MarshalIndent(*r, " ", " ")

	fmt.Fprintf(w, "Request=%s", ans)
}

func main() {
	// Przypisanie funkcji do ścieżki
	http.HandleFunc("/", handler)

	// Uruchomienie serwera
	http.ListenAndServe(":8080", nil)
}
