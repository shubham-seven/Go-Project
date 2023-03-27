package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fact int = 1

func main() {
	fmt.Println("Enter number for calculating its factorial:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input1 := strings.TrimSpace(input)
	n, err := strconv.Atoi(input1)
	if err != nil {
		panic(err)
	}

	res := Factorial(n)
	fmt.Println(res)

}
func Factorial(n int) int {

	if n >= 2 {
		fact = fact * n
		n--

		Factorial(n)
	}
	return fact
}
