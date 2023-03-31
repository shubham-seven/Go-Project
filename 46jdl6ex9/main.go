package main

import "fmt"

func foo(f func(x []int) int, ii []int) int {
	total := f(ii)

	return total
}

func main() {
	g := func(x []int) int {
		total := 0
		for _, v := range x {
			total = total + v

		}
		return total
	}
	ii := []int{
		2, 3, 4, 5, 8, 0, 9, 6, 7, 8,
	}

	res := foo(g, ii)
	fmt.Println(res)
}
