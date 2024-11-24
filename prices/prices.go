package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	fp, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("An error occured while opening prices.txt")
		panic(err)
	}

	scanner := bufio.NewScanner(fp)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("An error occured while reading prices.txt")
		fp.Close()
		panic(err)
	}

	prices := make([]float64, len(lines))

	for _, line := range lines {
		price, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("An error occured while parsing prices.txt")
			panic(err)
		}

		prices = append(prices, price)
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

}

func NewJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
