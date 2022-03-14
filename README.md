This package provides a function that takes in a float64 number and an int which specifies the amount of decimal spaces to round the float to.

Provide the float you want rounded and the number of decimal places you want it rounded to. The error is triggered if the provided int is <= 0.

Internally, it uses the math.Round() function provided by the math package in Go's standard library.