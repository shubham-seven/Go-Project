package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions class")
    greeter()
	result:=adder(3,5)
	fmt.Println(result)
	proResult,proMessage:=proAdder(2,3,6,9)
	fmt.Println(proResult)
	fmt.Println(proMessage)


}

func greeter(){
	fmt.Println("Hello From Go-Lang")
}

func adder(valOne int,valTwo int) int{
	return valOne + valTwo
}

func proAdder(Values...int)(int,string){
	total:=0
	for _,val:=range Values{
		
		//total=total+val
		total+=val

	}
    return total,"You have learned Pro functions"  

}