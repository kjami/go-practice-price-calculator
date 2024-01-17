package price_calculator

import (
	"errors"
	"os"
)

func ReadFile[T []float64 | []int | []string](fileName string, transform func([]byte) T) (T, error) {
	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("error while fetching file: " + fileName + " " + err.Error())
	}

	return transform(contentBytes), nil
}
