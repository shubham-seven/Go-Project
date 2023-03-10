package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop class")

	days := []string{"sunday","Monday","Tuesday","Wednesday","Thusrday","Friday","Saturday"}

	fmt.Println("Name of days in week are :")

	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}

	//Another Syntax for writting loop

	for i:= range days{
		fmt.Println(days[i])
	}
	// Another Syntax for writting loop
	for index,day:=range days{
		fmt.Printf("index is %v and day is %v\n",index,day)
	}
}