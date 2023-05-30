package calculator

import (
	"errors"
	"math"
)

func Plus(a, b int) int {
	return a + b
}

func PositivePow(a, n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("n is not positive")
	}
	return math.Pow(a, n), nil
}
