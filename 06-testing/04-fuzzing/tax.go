package tax

func CalculateTax(value float64) float64 {
	if value <= 0 {
		return 0.0
	} else if value >= 1000 && value < 20000 {
		return 10.0
	} else if value >= 20000 {
		return 20.0
	}
	return 5.0
}