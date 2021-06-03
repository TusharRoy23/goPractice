package main

import (
	"fmt"
	"strconv"
)

func concatFunc(n string) int {
	con := n
	sum := 0
	for i := 0; i < 3; i++ {
		v, err := strconv.Atoi(con)
		if err != nil {
			break
		}
		sum += v
		con += n
	}

	return sum
}

func main() {
	v1 := concatFunc("5")
	fmt.Println(v1)
	v2 := concatFunc("3")
	fmt.Println(v2)
}
