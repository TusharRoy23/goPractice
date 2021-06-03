package main

import (
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("abc.txt")
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()
}
