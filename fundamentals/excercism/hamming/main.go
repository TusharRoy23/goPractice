package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Type DNA:")
	var dnaOne string
	fmt.Scanf("%s\n", &dnaOne)

	fmt.Println("Type Daughter DNA:")
	var daughterDna string
	fmt.Scanf("%s\n", &daughterDna)

	var hamming int

	for index, value := range dnaOne {
		match, _ := regexp.MatchString("^[A-Z]+$", string(value))
		if match {
			if string(value) != string(daughterDna[index]) {
				hamming++
			}
		} else {
			fmt.Println("Special character found")
			hamming = 0
			break
		}
	}

	fmt.Println("Hamming Distance : ", hamming)
}
