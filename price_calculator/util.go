package price_calculator

import (
	"strconv"
	"strings"
)

func transformFileDataToFloat64Fn(data []byte) []float64 {
	strs := strings.Split(string(data), "\n")
	result := make([]float64, len(strs))
	for i, str := range strs {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			//ToDo: change this error handling
			panic(err)
		}
		result[i] = val
	}
	return result
}
