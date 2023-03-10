package main

import (
	"fmt"

)

func main() {
	fmt.Println("Welcome to class of Defer  : ")

defer myDefer()
defer fmt.Println("World")
defer fmt.Println("one")
defer fmt.Println("Two")
fmt.Println("Hello")

//defer myDefer()
}

func myDefer(){
	for i:=0;i<5;i++{
	defer fmt.Println(i)
	}
}