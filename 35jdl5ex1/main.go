package main

import "fmt"

type person struct {
	fname     string
	lname     string
	ficecream []string
}

func main() {

	p1 := person{
		fname: "shubham",
		lname: "trivedi",
		ficecream: []string{
			"butterscotch", "cappucino", "frooty",
		},
	}
	p2 := person{
		fname: "sumit",
		lname: "talreja",
		ficecream: []string{
			"strawberry", "vanila", "chocolate",
		},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.ficecream {
		fmt.Println(i, v)
	}
	for j, val := range p2.ficecream {
		fmt.Println(j, val)
	}
}
