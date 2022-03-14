package roundeasy

import (
	"errors"
	"math"
)

// Returns a rounded float64 to a specified number of decimal places
func Round(f float64, i int) (float64, error) {
	var err error = nil
	if i <= 0 {
		err = errors.New("decimal value lower or equal to 0. provide an int value greater than 0")
		return 0, err
	}

	rf := float64(i) * 10

	return math.Round(f*rf) / rf, err
}
