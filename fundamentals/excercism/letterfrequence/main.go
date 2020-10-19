package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type freqMap map[string]int

func frequency(s string) freqMap {
	m := freqMap{}
	for _, r := range s {
		m[string(r)]++
	}
	return m
}

func main() {
	fmt.Println("Type your sentence: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		sentence := scanner.Text()

		senSlice := strings.Fields(sentence)

		resChan := make(chan freqMap, len(senSlice))

		for _, s := range senSlice {
			go func(s string) {
				resChan <- frequency(s)
			}(s)
		}

		res := make(freqMap)
		for range senSlice {
			for index, freq := range <-resChan {
				res[index] += freq
			}
		}
		fmt.Println(res)
	}
}
