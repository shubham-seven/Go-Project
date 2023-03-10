package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	userinput := bufio.NewReader(os.Stdin)

fmt.Printf("How much rating will you give to pizza :")

//comma ok  // ERR ok

input,_ := userinput.ReadString('\n')
fmt.Println("Thanks for the rating",input)
fmt.Printf("Type of the string is: %T",input)
}