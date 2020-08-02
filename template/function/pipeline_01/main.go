package main

import (
	"os"
	"text/template"
	"log"
	"math"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl": double,
	"fsqr": square,
	"fsqrt": squareRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(float64(x))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 5)
	if err != nil {
		log.Fatalln(err)
	}
}