package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	for i, v := range myArray {
		fmt.Printf("O indice é %d e o seu valor é %d\n", i, v)
	}
}
