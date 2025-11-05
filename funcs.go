package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = x / 2
	for range 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	formated := fmt.Sprintf("Final approximation: %.2f", z)
	fmt.Println(formated)
	return z
}
