package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat (strings []string) ([]float64, error){
	var floats []float64

	for _, strings := range strings {
		floarPrice, err := strconv.ParseFloat(strings, 64)

		if err != nil {
			return nil, errors.New("Could not convert string to float64 !!!")
		}

	  	floats = append(floats, floarPrice)
	}
	return floats, nil
}