package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Funkcja wywoływana przez filepath.Walk dla każdego obiektu znalezionego w
// podanej ścieżce
func analiza(path string, info os.FileInfo, err error) error {
	fmt.Printf("Obiekt: %s=%+v\n", path, info)
	return nil
}

func main() {
	filepath.Walk("/usr/include", analiza)
}
