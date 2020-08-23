package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\n%b\n%#x\n", b, b, b)
}
