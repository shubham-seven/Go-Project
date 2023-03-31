package main

import "fmt"

var x int = 7
var g func() = func() {
	fmt.Println("this is g func")
}

func main() {
	//Anonymous func
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	//Assingning a function to a variable
	x := func() {
		fmt.Println("assigning a function to a variable")
	}
	x()
	g()

	g = x
	g()

	fmt.Println("done")
}
