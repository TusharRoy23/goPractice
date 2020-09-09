package main

import "fmt"

func main() {
	fun := func() string {
		return "abcd"
	}
	fmt.Println(fun())
}
