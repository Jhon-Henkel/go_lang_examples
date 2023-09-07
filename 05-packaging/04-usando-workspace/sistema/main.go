package main

// usar o go.work o comando é go work init ./math ./sistema
// isso pode ficar no .gitignore
// o go mod tidy não vai usar o workspace, ele vai usar o go.mod
// o comando go mod tidy -e vai ignorar os erros ao fazer o tidy
import "github.com/Jhon-Henkel/go_lang_examples/05-packaging/04-usando-workspace/math"

func main() {
	math2 := math.NewMath(1, 2)
	println(math2.Sum())
}