package main

import "fmt"

func swapFunc(ab, cd *float64) {
	*ab, *cd = *cd, *ab
}

func main() {
	x := 12.12
	fmt.Println("Address of x = ", &x)

	ptr := &x
	fmt.Printf("ptr type: %T, ptr value: %v\n", ptr, ptr)

	fmt.Printf("X value %v\n", *ptr)

	// 02 ex

	a, b := 10, 2

	ptra, ptrb := &a, &b

	fmt.Printf("ptra %v , ptrb %v\n", *ptra, *ptrb)

	e, f := 5.5, 8.8
	swapFunc(&e, &f)

	fmt.Printf("f %v , e %v\n", f, e)
}
