package main

import "math/rand"

func main() {
	a := rand.Int()
	b := rand.Int()

	if a > b {
		println("a is greater than b")
	} else {
		println("a is not greater than b")
	}

	switch a {
	case 5:
		println("equals 5")
	case 10:
		println("equals 10")
	default:
		println("neither 5 nor 10")
	}

	// para fazer o build para um os especifico basta rodar o comando: GOOS=windows go build
	// para mac intel: GOOS=darwin GOARCH=amd64 go build
	// go tool dist list esse comando mostra todos os sistemas operacionais e arquiteturas que o go suporta
}
