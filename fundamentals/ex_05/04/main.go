package main

import "fmt"

func main() {
	// map[keyType]valueType
	ano := struct {
		Name    string
		Age     int
		DevLang []string
		Numbers map[string]int
	}{
		Name: "Tushar",
		Age:  28,
		DevLang: []string{
			"Golang",
			"Javascript",
			"C++",
		},
		Numbers: map[string]int{
			"First":  13,
			"Second": 8,
			"Third":  12,
		},
	}

	fmt.Println(ano)
	fmt.Println(ano.Numbers["First"])
}
