package main

import "fmt"

func main() {
	i := 2
	if i == 3 {
		fmt.Println(3)
	} else if i == 2 {
		fmt.Println(2)
	} else {
		fmt.Println("Not found !")
	}
}
