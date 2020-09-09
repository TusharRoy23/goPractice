package main

import "fmt"

func main() {
	x := 5

	func(n int) {
		fmt.Println(n)
	}(x)
}
