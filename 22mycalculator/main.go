package main

import (
	"fmt"
	//"strings"
	// "math"
	//"bufio"
	"os"
	//"strconv"
)

func main() {
	fmt.Println("Welcome to my calculator")
	var float1,float2 float64
	fmt.Println("Enter first number:")

	_,err1:=fmt.Scanf("%f",&float1)
    CheckErrorNil(err1)
	fmt.Println("Enter Second number:")

	_,err2:=fmt.Scanf("%f",&float2)
	CheckErrorNil(err2)

	/*reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first number:")
	val1,_:=reader1.ReadString('\n')
	float1,err:=strconv.ParseFloat(strings.TrimSpace(val1),64)
    CheckErrorNil(err)


	reader2 := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Second number:")
	val2,_:=reader2.ReadString('\n')
    float2,err:=strconv.ParseFloat(strings.TrimSpace(val2),64)
    CheckErrorNil(err)*/

	Sum := float1 + float2
	fmt.Printf("Here is the Sum of two numbers: %f\n", Sum)

}
func CheckErrorNil(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
