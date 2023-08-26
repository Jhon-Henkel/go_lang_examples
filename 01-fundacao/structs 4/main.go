package main

import "fmt"

type Account struct {
	value int
}

func NewAccount(value int) *Account {
	return &Account{value: value}
}

func (a Account) simulate(value int) int {
	a.value += value
	fmt.Printf("O valor simulado é %d\n", a.value)
	return value
}

func (a *Account) approve(value int) int {
	a.value += value
	fmt.Printf("O saldo após aprovar é %d\n", a.value)
	return a.value
}

func main() {
	account := Account{value: 1000}
	simulated := account.simulate(500)
	account.approve(simulated)
	fmt.Printf("O saldo final é %d\n", account.value)
}
