package cmdmanager

import "fmt"

type CMDManager struct{}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter your prices:")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		_, err := fmt.Scanln(&price)

		if err != nil {
			return nil, err
		}

		if price == "" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
