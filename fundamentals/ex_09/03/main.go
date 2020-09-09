package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of cpus: ", runtime.NumCPU())
	fmt.Println("Number of go routines: ", runtime.NumGoroutine())
	inc := 0
	goSch := 100
	var wg sync.WaitGroup
	wg.Add(goSch)

	for i := 0; i <= goSch; i++ {
		go func() {
			v := inc
			runtime.Gosched()
			v++
			inc = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", inc)
}
