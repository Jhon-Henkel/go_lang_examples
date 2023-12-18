package main

import "fmt"

// Não é obrigatório a seta (dizer a direção do canal), mas é uma boa prática
func receiver(name string, data chan<- string) { // essa seta significa que o canal vai apenas receber informação
	data <- name
}

func reader(data <-chan string) { // essa seta significa que o canal vai apenas esvaziar informação
	fmt.Println(<-data)
}

func main() {
	data := make(chan string)
	go receiver("Hello World!", data)
	reader(data)
}
