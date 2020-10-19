package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	selectedTime := time.Now()
	power := time.Duration(math.Pow(10, 9)) * time.Second // type conversion => time.Duration()

	plus := selectedTime.Add(power)

	fmt.Println(plus)
}
