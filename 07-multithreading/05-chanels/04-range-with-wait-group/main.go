package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	chanel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	go publish(chanel)
	go reader(chanel, &waitGroup)
	waitGroup.Wait()
}

func reader(chanel chan int, waitGroup *sync.WaitGroup) {
	for x := range chanel {
		fmt.Printf("Received %d\n", x)
		waitGroup.Done()
	}
}

func publish(chanel chan int) {
	for i := 0; i < 10; i++ {
		chanel <- i
	}
	close(chanel) // indicando que nada mais vai entrar no canal
}
