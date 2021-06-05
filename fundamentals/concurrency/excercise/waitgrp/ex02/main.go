package main

import (
	"fmt"
	"sync"
)

func sum(a float64, b float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f\n", a+b)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sum(2.5, 3.6, &wg)
	go sum(2.4, 3.4, &wg)
	go sum(2.2, 3.2, &wg)

	wg.Wait()
}
