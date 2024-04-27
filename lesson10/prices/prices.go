package prices

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process()  {
	result := make(map[string]float64, 0)

		for priceIndex, price := range job.InputPrices {
			result[price] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
