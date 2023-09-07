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

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expectedTax float64
	}

	table := []calcTax{
		{100, 5},
		{500, 5},
		{1000, 10},
		{2000, 10},
		{3000, 10},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)
		if result != test.expectedTax {
			t.Errorf("Result %.2f, expected %.2f", result, test.expectedTax)
		}
	}
}
