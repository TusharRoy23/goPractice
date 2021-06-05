package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	carInfo := car{licenceNo: "acedf", brand: "BMW"}
	fmt.Println(carInfo.License())
	fmt.Println(carInfo.Name())
}
