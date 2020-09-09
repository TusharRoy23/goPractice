package main

import "fmt"

func main() {
	funcVariable := func(param int) int {
		var bar func(recursiveParam int) int
		bar = func(recursiveParam int) int {
			if recursiveParam == 0 {
				return 1
			}
			return recursiveParam * bar(recursiveParam-1)
		}
		return bar(param)
	}
	argument := 4

	recursive := foo(funcVariable, argument)
	fmt.Println(recursive)
}

func foo(f func(needParam int) int, param int) int { // callback
	return f(param)
}
