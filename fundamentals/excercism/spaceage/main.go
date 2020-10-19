package main

import (
	"fmt"
)

func main() {
	// earthYr := float64(1.0)
	// earthDay := float64(365.25)
	earthSec := float64(31557600)

	fmt.Println("Tell me the seconds:")

	var sec int
	_, err := fmt.Scanf("%d", &sec)

	if err != nil {
		fmt.Println(err)
	}

	// input_sec / ( earth_sec * self_sec )
	mercury := float64(earthSec * 0.2408467)
	venus := float64(earthSec * 0.61519726)
	mars := float64(earthSec * 1.8808158)
	jupiter := float64(earthSec * 11.862615)
	saturn := float64(earthSec * 29.447498)
	uranus := float64(earthSec * 84.016846)
	neptune := float64(earthSec * 164.79132)

	fmt.Printf("age in mercury %.2f yrs\n", float64(sec)/mercury)
	fmt.Printf("age in venus %.2f yrs\n", float64(sec)/venus)
	fmt.Printf("age in mars %.2f yrs\n", float64(sec)/mars)
	fmt.Printf("age in jupiter %.2f yrs\n", float64(sec)/jupiter)
	fmt.Printf("age in saturn %.2f yrs\n", float64(sec)/saturn)
	fmt.Printf("age in uranus %.2f yrs\n", float64(sec)/uranus)
	fmt.Printf("age in neptune %.2f yrs\n", float64(sec)/neptune)
	fmt.Printf("age in earth %.2f yrs\n", float64(sec)/earthSec)
}
