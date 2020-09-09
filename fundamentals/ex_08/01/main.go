package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{"Tushar", 28}
	p2 := person{"Reja", 25}
	p3 := person{"Rahsut", 27}

	persons := []person{p1, p2, p3}

	byteSlice, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteSlice))

	var people []person

	err = json.Unmarshal(byteSlice, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", people)
}
