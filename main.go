package main

import (
	"example.com/priceCalculator/cmdmanager"
	"example.com/priceCalculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmd := cmdmanager.New()
		priceJob := prices.NewJob(cmd, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process price job")
			fmt.Println(err)
		}
	}

}
