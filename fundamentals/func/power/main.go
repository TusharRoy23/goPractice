package main

import (
	"fmt"
	"math"
)

func power(a float64) float64 {
	return math.Pow(a, 3)
}

func main() {
	v := power(3.)
	fmt.Printf("%v\n", v)
}
