package main

import "fmt"

func main() {
	f1 := foo()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

	fmt.Println("----------")

	f2 := foo()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

	fmt.Println("----------")

	f3 := bar()
	fmt.Println(f3)
	fmt.Println(f3)
}

func foo() func() int { // func() int return type
	x := 0
	return func() int {
		x++
		return x
	}
}

func bar() int {
	x := 0
	x++
	return x
}
