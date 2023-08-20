package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
}

type Cliente struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func main() {
	jhon := Cliente{
		Name:   "Jhon",
		Age:    20,
		Active: true,
	}
	jhon.Active = false
	jhon.Address.Street = "Rua dos bobos"
	jhon.Address.Number = 10
	jhon.Address.City = "São Paulo"

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", jhon.Name, jhon.Age, jhon.Active)
	fmt.Printf("Rua: %s\nNúmero: %d\nCidade: %s\n", jhon.Address.Street, jhon.Address.Number, jhon.Address.City)
}
