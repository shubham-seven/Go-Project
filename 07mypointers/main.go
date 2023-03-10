package main

import "fmt"

func main(){

	fmt.Println("Welcome to class of pointers")

	var ptr *int 
	nmbr := 15
	ptr = &nmbr
	fmt.Println(ptr)
	fmt.Println(*ptr)

*ptr = *ptr +10
fmt.Println("New value of number is: ",nmbr)


}