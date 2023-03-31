package main

import "fmt"

type person struct{
value int
}
func changeMe(p *person){

*p=person{
	value:43,
}
fmt.Println(*p)
}

func main(){
p:= person{
	value:12,
}
changeMe(&p)
}