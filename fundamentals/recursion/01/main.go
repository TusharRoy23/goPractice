package main

import "fmt"

func reverse(name string, size int) string {
	if size == 0 {
		return string(name[0])
	}
	return string(name[size]) + reverse(name, (size-1))
}

func main() {
	name := "Tushar Roy Chowdhury"
	result := reverse(name, len(name)-1)
	fmt.Println(result)
}
