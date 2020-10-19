package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Type a sentence: ")

	inp := bufio.NewReader(os.Stdin)

	sentence, _ := inp.ReadString('\n')

	sliceSen := strings.Fields(sentence)

	fmt.Print("Acronym: ")

	fmt.Println(acronym(sliceSen))
}

func acronym(sliceSen []string) string {
	var str string
	for _, value := range sliceSen {
		match, _ := regexp.MatchString("^[a-zA-Z]+$", string(value[0]))

		if match {
			str = str + strings.ToUpper(string(value[0]))
		} else {
			return "Only Alphabets"
		}
	}
	return str
}
