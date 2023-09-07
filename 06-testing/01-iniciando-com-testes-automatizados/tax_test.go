package tax

import "testing"

// para rodar o teste, basta usar o comando go test . dentro da pasta em que e encontra o teste
// o comando go test -v . vai mostrar o resultado de cada teste
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedTax := 6.0

	result := CalculateTax(amount)

	if result != expectedTax {
		t.Errorf("Result %.2f, expected %.2f", result, expectedTax)
	}
}