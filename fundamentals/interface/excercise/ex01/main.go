package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {
	var salary money = 1.5 * 1.253
	salary.print()

	s := salary.printStr()
	fmt.Println(s)
}
