package main

func main() {
	// Deve-se evitar usar isso, mas para usar o buffer é só passar o tamanho do buffer no segundo parâmetro
	// Não adianta utilizar um valor exorbitante, pois o buffer isso só vai fazer com que o programa fique mais lento
	chanel := make(chan string, 2)
	chanel <- "Hello"
	chanel <- "World"

	println(<-chanel)
	println(<-chanel)
}
