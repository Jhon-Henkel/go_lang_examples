package main

import "time"

func main() {
	chanelOne := make(chan int)
	chanelTwo := make(chan int)

	go func() {
		time.Sleep(time.Second)
		chanelOne <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		chanelTwo <- 2
	}()

	// Se quiser pegar informações de fila, por exemplo, pode usar um for, gerando um loop infinito
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chanelOne:
			println("received", msg1)
		case msg2 := <-chanelTwo:
			println("received", msg2)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
