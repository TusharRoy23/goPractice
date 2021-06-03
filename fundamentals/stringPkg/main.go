package main

import "fmt"

func main() {
	// Exercise 01

	/* var name string = "Tushar"
	country := "Bangladesh"

	fmt.Println(name)
	fmt.Println(country)

	fmt.Println(`He says: "Hello"`)
	fmt.Println(`C:\Users\a.txt`) */

	// Excerxise 02

	/* r := "ă"

	fmt.Printf("%T\n", r)

	var str1 string = "ma"
	var str2 string = "m"

	str := str1 + str2 + string(r)
	fmt.Printf("str is %s\n", str) */

	s := "你好 Go!"
	b := []rune(s)

	fmt.Printf("%#v\n", b)
	for i, v := range b {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
