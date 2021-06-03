package main

import (
	"log"
	"os"
)

func main() {
	// check file exists or not
	_, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exists")
		}
	}
	// renaming the file
	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal(err)
	}

}
