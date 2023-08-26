package main

import (
	"fmt"
)

// O defer é um statement que adia a execução de uma função até o final do escopo
// Serve para fechar conexões, arquivos, etc, pois o defer é executado depois de tudo
func main() {
	defer fmt.Println("primeira linha")
	fmt.Println("segunda linha")
	fmt.Println("terceira linha")
}
