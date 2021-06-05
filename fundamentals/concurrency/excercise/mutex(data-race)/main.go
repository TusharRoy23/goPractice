package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mututal *sync.Mutex) {
	mututal.Lock()
	*b += n
	mututal.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mututal *sync.Mutex) {
	mututal.Lock()
	*b -= n
	mututal.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mututal sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mututal)
		go withdraw(&balance, i, &wg, &mututal)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
