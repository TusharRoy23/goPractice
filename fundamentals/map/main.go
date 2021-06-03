package main

import (
	"fmt"
)

func main() {
	var m1 map[string]string
	fmt.Printf("%v value, %T Type\n ", m1, m1)

	m2 := map[int]string{
		10: "Abba",
	}

	fmt.Println("Existing", m2[10])
	fmt.Println("Non existing", m2[1])
}
