package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Type valid number: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		number := scanner.Text()

		match, _ := regexp.MatchString("^[0-9 ]+$", number)

		if match {
			var totalNum int64 = 0
			sliceNum := strings.Fields(number)

			for i := len(sliceNum) - 1; i >= 0; i++ {
				len := len(sliceNum[i])
				for len > 0 {
					if i%2 != 0 {
						num, _ := strconv.ParseInt(string(sliceNum[i][len-1]), 36, 32)
						num = num * 2
						if num > 9 {
							num = num - 9
						}
						totalNum += num
					} else {
						num, _ := strconv.ParseInt(string(sliceNum[i][len-1]), 36, 32)
						totalNum += num
					}
					len--
				}
			}

			even := totalNum / 10
			if even%2 == 0 {
				fmt.Println("Valid Number")
			} else {
				fmt.Println("Invalid Number")
			}
		} else {
			fmt.Println("Only digit & space are allowed")
		}
	}
}
