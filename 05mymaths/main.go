package main

import (
    "math/rand"
	"fmt"
	//"crypto/rand"
	)

func main(){

	var mathnumber1 int = 2
	var mathnumber2 float64 = 4

	fmt.Println("sum of these two numbers  are :",mathnumber1 + int(mathnumber2))

	//Random Number

	rand.Seed(7)
	fmt.Println(rand.Intn(5))



}