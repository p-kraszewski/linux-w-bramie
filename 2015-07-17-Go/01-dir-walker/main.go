package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func analiza(path string, info os.FileInfo, err error) error {
	fmt.Printf("Obiekt: %s\n", info.Name())
	return nil
}

func main() {
	filepath.Walk("/usr/include", analiza)
}
