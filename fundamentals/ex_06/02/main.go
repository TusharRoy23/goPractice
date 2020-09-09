package main

import "fmt"

func main() {
	defer foo()
	defer test()
	bar()
	/*
		Defer function works from bottom to top,
		test() -> foo()
	*/
}

func foo() {
	fmt.Println("I am foo")
}

func bar() {
	fmt.Println("I am bar")
}

func test() {
	fmt.Println("I am test")
}
