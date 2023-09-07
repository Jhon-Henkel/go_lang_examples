package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedTax := 6.0

	result := CalculateTax(amount)

	if result != expectedTax {
		t.Errorf("Result %.2f, expected %.2f", result, expectedTax)
	}
}

// comando para benchmarking:
// go test -bench=.

// BenchmarkCalculateTax-4       -  1000000000              -  0.3458 ns/op    			
// esse -4 é a quantidade de cpu -  quantidade de operaçoes - operações por nanosegundo

// go help test - para mais informações de comandos de testes

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
