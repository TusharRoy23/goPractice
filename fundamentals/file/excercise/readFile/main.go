package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	data, err := ioutil.ReadAll(newFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
}
