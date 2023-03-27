package main

import "fmt"

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func even(f func(x ...int) int, v ...int) int {
	var y []int
	for _, i := range v {
		if i%2 == 0 {
			y = append(y, i)

		}
	}
	return f(y...)
}
func odd(f func(x ...int) int, v ...int) int {
	var y []int
	for _, i := range v {
		if i%2 != 0 {
			y = append(y, i)

		}
	}
	return f(y...)
}

func main() {
	slice1 := []int{2, 3, 5, 8, 9, 0, 7, 5, 3, 2, 1}
	res := sum(slice1...)
	fmt.Println(res)

	res2 := even(sum, slice1...)
	fmt.Println(res2)
	res3 := odd(sum, slice1...)
	fmt.Println(res3)

}
