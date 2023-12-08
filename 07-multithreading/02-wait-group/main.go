package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		waitGroup.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	// O número 25 é a quantidade de processos que serão executados em paralelo
	waitGroup.Add(25)
	go task("A", &waitGroup)
	go task("B", &waitGroup)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d Task %s is running\n", i, "anonimous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
