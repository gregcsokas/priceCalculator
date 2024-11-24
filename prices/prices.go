package prices

import (
	"example.com/priceCalculator/conversion"
	"example.com/priceCalculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IoManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {

	lines, err := job.IoManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadPrices()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IoManager.WriteResult(job)

}

func NewJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IoManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
