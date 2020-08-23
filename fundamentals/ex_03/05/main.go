package main

import "fmt"

func main() {
	i := 10
	for i <= 100 {
		fmt.Printf("%v\n", i%4)
		i++
	}
}
