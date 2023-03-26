package main

import "fmt"

func main() {
	/*
		for i:=1;i<=1000;i++{
			fmt.Println(i)
		}
	*/
	i := 1
	for j := 65; j <= 90; j++ {

		fmt.Println("\n", i)
		i++
		for k := 0; k < 3; k++ {
			fmt.Printf("%#U\t\n", j)
		}
	}
}
