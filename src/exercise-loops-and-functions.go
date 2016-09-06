package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	for {
		newZ := z - (z*z-x)/(2*z)
		if math.Abs(newZ - z) < 1e-8 {
			break
		} else {
			z = newZ
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
