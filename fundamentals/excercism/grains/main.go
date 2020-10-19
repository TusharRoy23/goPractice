package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("How many grains as start: ")
	var grainStart int
	var totalGrain uint64

	_, err := fmt.Scanf("%d\n", &grainStart)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 64; i++ {
		value, err := square(grainStart)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print((i + 1), " Number of index: ")
		fmt.Printf("%d\n", value)
		grainStart++
	}

	totalGrain = 1<<64 - 2
	fmt.Println("Total grain: ", totalGrain)
}

func square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Selected number must be between 1 and 64")
	}
	return 1 << (n - 1), nil
}
