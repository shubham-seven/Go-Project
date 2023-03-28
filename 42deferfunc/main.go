package main

import "fmt"

func main(){

add()

value:=person{first : "shubham",last : "trivedi", age: 31,}
value.speak()

}
func add (){
	defer dfr()
	fmt.Println("add function:") 
	
}
 func dfr(){
	fmt.Println("This is deferred function")
	
}
func (p person)speak(){
fmt.Println("Your name: ",p.first," ",p.last," Your age: ",p.age)
}
type person struct{
	first string
	last string
	age int

}