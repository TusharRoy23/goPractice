package main

import (
	"fmt"
	"strings"
)

func searchItem(search string, strSlic ...string) bool {
	statement := false
	search = strings.ToLower(search)
	for _, value := range strSlic {
		if value == search {
			statement = true
			break
		}
	}
	return statement
}

func main() {
	result := searchItem("HORSE", "lion", "tiger", "bear")
	fmt.Println(result)
}
