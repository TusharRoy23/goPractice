package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sum, product := 0, 1

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")
	}
	args := os.Args[1:]
	for _, value := range args {
		num, err := strconv.ParseInt(value, 0, 32)
		if err != nil {
			continue
		}
		sum += int(num)
		product *= int(num)
	}

	fmt.Println(sum)
	fmt.Println(product)
}
