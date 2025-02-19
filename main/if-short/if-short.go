package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim { // V initialized with Pow(x,y) and then v is checked if it's lowet than lim
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(2, 4, 10))
}
