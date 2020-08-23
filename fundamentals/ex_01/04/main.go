package main

import "fmt"

type hotdog int // custom type but underlying int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
