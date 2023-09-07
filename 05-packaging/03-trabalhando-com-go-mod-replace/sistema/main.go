package main

// como esse pacote ainda não existe no github, o go mod replace vai substituir o pacote
// o comando é go mod edit -replace github.com/Jhon-Henkel/go_lang_examples/05-packaging/03-trabalhando-com-go-mod-replace/math=../math
// após isso o comando go mod tidy vai atualizar o go.mod
// NÃO ESQUECER DE REMOVER O REPLACE DO GO MOD antes de publicar o pacote, isso é gambiarra, o certo é usar a forma da pasta 05-packaging/04-usando-workspace
import "github.com/Jhon-Henkel/go_lang_examples/05-packaging/03-trabalhando-com-go-mod-replace/math"

func main() {
	math2 := math.NewMath(1, 2)
	println(math2.Sum())
}