package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(n float64) {
		fmt.Printf("%f\n", math.Sqrt(n))
		wg.Done()
	}(30.33)
	wg.Wait()
}
