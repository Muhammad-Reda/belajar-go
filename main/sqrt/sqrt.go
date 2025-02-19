package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	counter := 0
	z := x
	var temp float64
	for v := 1; v <= 10; v++ {
		temp = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z, temp)
		if z == temp {
			return z, counter
		}
		counter++
	}
	return z, counter
}

func main() {
	fmt.Println(sqrt(4))
	fmt.Println(math.Sqrt(4))
}
