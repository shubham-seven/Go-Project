package main

import "fmt"

type person struct{
	first string
	last string
}
type secretAgent struct{
	person
	ltk bool
}
func (p person) Speak(){
fmt.Println("I am ",p.first," ",p.last)
}
func (s secretAgent) Speak(){
	fmt.Println("I am ",s.first," ",s.last," has licence to kill : ",s.ltk)
	}
type human interface {
	Speak()
}
func  bar(h human){
	switch h.(type) {
	case person :
		fmt.Println("I was passed into the bar ",h.(person).first," ",h.(person).last)
		case secretAgent : 
		fmt.Println("I was passed into the bar ",h.(secretAgent).first,h.(secretAgent).last,h.(secretAgent).ltk)
	}
fmt.Println("I was passed into the bar ",h)
}
func main(){

	sa1 := secretAgent{
		person : person{first:"Shubham",last:"Trivedi",},
		ltk : true,
	}
   sa2 := secretAgent{
	person : person{first:"Sumit",last:"Talreja",},
	ltk : true,
}
p1 := person{
	first: "land",last: "lord",
}
sa1.Speak()
sa2.Speak()
p1.Speak()
bar(sa1)
bar(sa2)
bar(p1)
}