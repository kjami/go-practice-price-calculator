package price_calculator

import (
	"fmt"
)

type TaxCalculationsJob struct {
	TaxRate       float64
	Prices        []float64
	PricesWithTax map[string]float64
}

func NewTaxCalculationsJob(taxRate float64) *TaxCalculationsJob {
	fileData, err := ReadFile("price_calculator/prices.txt", transformFileDataToFloat64Fn)

	if err != nil {
		panic(err)
	}
	return &TaxCalculationsJob{
		TaxRate: taxRate,
		Prices:  fileData,
		//PricesWithTax will be empty for now
	}
}

func (job *TaxCalculationsJob) Process() {
	result := make(map[string]float64, len(job.Prices))
	taxFn := getTaxCalculatorFn(job.TaxRate)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = taxFn(price)
	}
	job.PricesWithTax = result
	fmt.Println(result)
}

func getTaxCalculatorFn(tax float64) func(float64) float64 {
	return func(number float64) float64 {
		return number * (1 + tax)
	}
}
