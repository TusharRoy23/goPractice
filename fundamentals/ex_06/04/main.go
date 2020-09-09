package main

import (
	"fmt"
	"math"
)

type square struct {
	Length float64
}

type circle struct {
	Radius float64
}

func (sq square) area() float64 {
	return sq.Length * sq.Length
}

func (cr circle) area() float64 {
	return math.Pi * cr.Radius * cr.Radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		Length: 2,
	}
	// sqVal := sq.area()
	// fmt.Println(sqVal)

	cr := circle{
		Radius: 12.35,
	}
	// crVal := cr.area()
	// fmt.Println(crVal)
	info(sq)
	info(cr)
}
