package main

import "fmt"

type Hotdog int

var x Hotdog

func main(){

	fmt.Println(x)
	fmt.Printf("Type is:%T\n",x)
 x = 42
 fmt.Println(x)

}