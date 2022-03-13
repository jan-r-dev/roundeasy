This package provides two functions:

round32(float32, int) float32, err
round64(float64, int) float64, err

Provide the float you want rounded and the number of decimal places you want it rounded to. The error is triggered if the provided int is <= 0.

Internally, it uses the math.Round() function provided by the math package in Go's standard library.