package main

import "fmt"

type Cliente struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	jhon := Cliente{
		Name:   "Jhon",
		Age:    20,
		Active: true,
	}
	jhon.Active = false

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", jhon.Name, jhon.Age, jhon.Active)
}
