package main

import (
	"fmt"
	"math"
	"sort"
)

type kind int

const (
	NaT kind = iota
	Equ
	Iso
	Sca
)

func main() {
	var angOne float64 = 0
	var angTwo float64 = 0
	var angThree float64 = 0

	fmt.Println("Angle One: ")
	fmt.Scanf("%f", &angOne)

	fmt.Println("Angle Two: ")
	fmt.Scanf("%f", &angTwo)

	fmt.Println("Angle Three: ")
	fmt.Scanf("%f", &angThree)

	fmt.Println(sideCheck(angOne, angTwo, angThree))
}

func sideCheck(angOne, angTwo, angThree float64) kind {
	sides := [3]float64{angOne, angTwo, angThree}
	sort.Float64s(sides[:])

	for _, f := range sides {
		if f <= 0 || math.IsNaN(f) || math.IsInf(f, 0) {
			return NaT
		}
	}

	if sides[0]+sides[1] < sides[2] {
		return NaT
	}

	if sides[0] == sides[1] && sides[0] == sides[2] {
		return Equ
	}

	if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}

	return Sca
}
