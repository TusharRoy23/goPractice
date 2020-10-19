package main

import (
	"fmt"
)

func main() {
	fmt.Println("Type a name: ")
	var name string
	fmt.Scanf("%s\n", &name)

	fmt.Println(twoFer(name))
}

func twoFer(name string) string {
	if name != "" {
		return "One for " + name + " One for me"
	}
	return "One for you and One for me"
}
