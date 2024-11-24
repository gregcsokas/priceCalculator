package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		price, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error converting %s to float", stringVal))
		}

		floats = append(floats, price)
	}

	return floats, nil
}
