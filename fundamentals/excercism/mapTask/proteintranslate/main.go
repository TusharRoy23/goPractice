package main

import (
	"bufio"
	"fmt"
	"os"
)

var protein map[string]string

func main() {
	fmt.Println("Type RNA sequence:")
	protein = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		sentence := scanner.Text()

		if len(sentence)%3 == 0 {
			str := ""
			ret := ""
			for index, val := range sentence {
				if index != 0 && (index+1)%3 == 0 {
					str += string(val)
					ret = getProtein(str)
					if ret != "" {
						fmt.Println(str, " : ", ret)
					} else {
						break
					}

					str = ""
				} else {
					str += string(val)
				}
			}
		} else {
			fmt.Println("This is not RNA sequence")
		}
	}
}

func getProtein(str string) string {
	for ind, val := range protein {
		// fmt.Println(str)
		if ind == str {
			if val == "STOP" {
				break
			} else {
				return val
			}
		}
	}
	return ""
}
