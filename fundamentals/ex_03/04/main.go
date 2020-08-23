package main

import "fmt"

func main() {
	b := 1992
	for {
		if b > 2020 {
			break
		}
		fmt.Println(b)
		b++
	}
}
