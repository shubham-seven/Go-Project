package main

import "fmt"

func main() {

	ii := []int{2, 6, 8, 9, 0, 3, 2, 1, 5, 6, 9}
	n := foo(ii...)
	fmt.Println(n)
	ii2 := []int{2, 6, 8, 9, 0, 3, 2, 1, 5, 6, 9, 10, 7, 18, 90}
	m := bar(ii2)

	fmt.Println(m)

}
func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func bar(y []int) int {
	var total int
	for _, v := range y {
		total += v
	}
	return total
}
