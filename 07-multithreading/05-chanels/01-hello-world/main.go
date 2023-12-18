package main

import "fmt"

// Thread 1
func main() {
	chanel := make(chan string) // canal está vazio

	// Thread 2
	go func() {
		chanel <- "Hello World!" // canal está cheio
	}()

	// Thread 1
	msg := <-chanel // canal está vazio
	fmt.Println(msg)
}
