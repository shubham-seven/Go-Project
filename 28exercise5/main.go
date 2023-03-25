package main

import "fmt"

type HotDog int

var x HotDog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y := int64(x)
	//y:=int32(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
