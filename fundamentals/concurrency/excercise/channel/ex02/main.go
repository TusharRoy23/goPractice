package main

import (
	"fmt"
	"math"
	"strings"
)

func power(n float64, c chan int) {
	c <- int(math.Pow(n, 2))
}

func main() {
	ch := make(chan int)
	for i := 1.; i <= 50.; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("%d\n", <-ch)
	}

	fmt.Println(strings.Repeat("#", 10))
	cha := make(chan int)

	for i := 1.; i <= 50.; i++ {
		go func(n float64) {
			cha <- int(math.Pow(n, 2))
		}(i)
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("%d\n", <-cha)
	}
}
