package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Say something to bob: ")

	inp := bufio.NewReader(os.Stdin)
	speech, _ := inp.ReadString('\n')

	speechSlice := strings.Fields(speech)
	length := len(speechSlice)

	if length == 0 {
		fmt.Println("Fine. Be that way!'")
	} else if string(speechSlice[length-1]) == "?" && checkUppercase(speech) {
		fmt.Println("Calm down, I know what I'm doing !")
	} else if string(speechSlice[length-1]) == "?" && !checkUppercase(speech) {
		fmt.Println("Sure !")
	} else if checkUppercase(speech) {
		fmt.Println("Whoa, chill out!")
	} else {
		fmt.Println("Whatever.")
	}
}

func checkUppercase(speech string) bool {
	if strings.ToUpper(speech) == speech {
		return true
	}
	return false
}
