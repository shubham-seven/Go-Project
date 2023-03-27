package main

import "fmt"

func main() {
	val1 := truck{
		vehicle: vehicle{doors: "4", colors: "red"}, fourWheel: false,
	}
	val2 := sedan{
		vehicle: vehicle{doors: "2", colors: "black"}, luxury: true,
	}
	fmt.Println(val1)
	fmt.Println(val2)

	fmt.Println(val1.vehicle.colors)
	fmt.Println(val1.vehicle.doors)
	fmt.Println(val1.fourWheel)

	fmt.Println(val2.vehicle.colors)
	fmt.Println(val2.vehicle.doors)
	fmt.Println(val2.luxury)

}

type vehicle struct {
	doors  string
	colors string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
