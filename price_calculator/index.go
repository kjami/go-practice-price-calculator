package price_calculator

func PriceCalculator() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		job := NewTaxCalculationsJob(taxRate)
		job.Process()
	}
}
