package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100.; i < 150.; i++ {
		go func(n float64, wg *sync.WaitGroup) {
			fmt.Printf("%.2f\n", math.Sqrt(n))
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
