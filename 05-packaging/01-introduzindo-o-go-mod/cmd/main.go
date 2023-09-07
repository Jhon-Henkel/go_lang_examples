package main

import (
	"fmt"
	"github.com/Jhon-Henkel/go_lang_examples/05-packaging/01-introduzindo-o-go-mod/math"
)

func main() {
	math := math.NewMath(10, 20)
	fmt.Println(math.Sum())
}
