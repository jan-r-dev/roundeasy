package roundeasy

import (
	"errors"
	"math"
)

func round32(f32 float32, decI int) (float32, error) {
	var err error = nil
	if decI <= 0 {
		err = errors.New("decimal value lower or equal to 0. provide an int value greater than 0")
		return 0, err
	}

	decF := float64(decI) * 10
	f64 := float64(f32)

	return float32(math.Round(f64*decF) / decF), err
}

func round64(f64 float64, decI int) (float64, error) {
	var err error = nil
	if decI <= 0 {
		err = errors.New("decimal value lower or equal to 0. provide an int value greater than 0")
		return 0, err
	}

	decF := float64(decI) * 10

	return (math.Round(f64) * decF) / decF, err
}

