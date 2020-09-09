package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "My name is Tushar Roy Chowdhury"
	// encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	s64 := base64.StdEncoding.EncodeToString([]byte(s)) // StdEncoding is equivalent to StdEncoding

	fmt.Println(len(s))
	fmt.Println(len(s64))

	fmt.Println(s)
	fmt.Println(s64)
}
