package main

import "fmt"

func X() func() {
	return func() {
		fmt.Println("I am a returned function from a function")
	}
}
func main() {
	x := func() {
		fmt.Println("I am a returned function from a function")

	}
	x()
	//Assingning a function to a variable
	y := X()
	y()
}
