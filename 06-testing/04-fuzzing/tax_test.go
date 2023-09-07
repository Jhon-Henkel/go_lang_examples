package tax

import "testing"

// fuzzing é um teste que gera dados aleatórios para testar o código
// para rodar o fuzz, execute o comando:
// go test -v -fuzz=.

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-5, 0, 500.0, 100.0, 1000.0, 2000.0, 3000.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Result %.2f, expected 0", result)
		}
		if amount >= 20000 && result != 20 {
			t.Errorf("Result %.2f, expected 20", result)
		}
	})
}
