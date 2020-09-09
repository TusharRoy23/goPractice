package main

import "fmt"

type vehicle struct {
	Doors int
	Color string
}

type truck struct {
	vehicle
	FourWheel bool
}

type sedan struct {
	vehicle
	Luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			Doors: 2,
			Color: "Yellow",
		},
		FourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			Doors: 4,
			Color: "Black",
		},
		Luxury: true,
	}

	fmt.Println("truck: ", t, " sedan: ", s)
	fmt.Println(t.vehicle.Doors)
	fmt.Println(s.vehicle.Doors)
}
