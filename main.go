package main

import (
	"example.com/priceCalculator/filemanager"
	"example.com/priceCalculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewJob(fm, taxRate)
		priceJob.Process()
	}

}
