package main

import (
	"fmt"
	"github.com/Jhon-Henkel/go-lang-full-cycle-secret-repository/pkg/events"
)

func main() {
	eventDispatcher := events.NewEventDispatcher()
	fmt.Println(eventDispatcher)
}
