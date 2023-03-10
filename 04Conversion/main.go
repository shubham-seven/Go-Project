package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	fmt.Println("Welcome to the Pizza App:")

	fmt.Println("Please rate our Pizza between 1 and 5 ")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for the Rating : ",input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Adding 1 to your Rating : ",numRating+1)
	}

	






}