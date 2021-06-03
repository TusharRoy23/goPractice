package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
