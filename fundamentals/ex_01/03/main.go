package main

import "fmt"

var x int = 2
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%d\n%s\n%t\n", x, y, z)
	fmt.Print(s)
}
