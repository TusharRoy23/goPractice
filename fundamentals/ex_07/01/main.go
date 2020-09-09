package main

import "fmt"

func main() {
	value := 8
	fmt.Println(&value)
	value2 := &value
	fmt.Println(*value2)
}
