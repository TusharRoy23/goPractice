package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var values map[string]int

func main() {
	values = map[string]int{
		"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
		"D": 2, "G": 2,
		"B": 3, "C": 3, "M": 3, "P": 3,
		"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
		"K": 5,
		"J": 8, "X": 8,
		"Q": 10, "Z": 10,
	}

	fmt.Println("Type word or sentence: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		sentence := scanner.Text()

		senSlice := strings.Fields(sentence)
		total := 0

		for _, val := range senSlice {
			total += getValue(string(val))
		}

		fmt.Println("Total: ", total)
	}
}

func getValue(slice string) int {
	total := 0
	for _, val := range slice {
		total += search(string(val))
	}
	return total
}

func search(str string) int {
	for key, val := range values {
		if strings.ToUpper(str) == key {
			return val
		}
	}
	return 0
}
