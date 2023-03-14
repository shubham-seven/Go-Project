package main
import (
	"fmt"
    "strings"
)

func main(){
	fmt.Println("Welcome to my advance calculator:")
	var val1,val2 float64
	var operation string
	fmt.Println("Enter the first number:")

	_,err1 := fmt.Scanf("%f",&val1)
	CheckErrorNil(err1)
	fmt.Println("Enter the Second number:")

	_,err2 := fmt.Scanf("%f",&val2)
	CheckErrorNil(err2)

	fmt.Println("Enter the operation you want to apply on numbers,select one operation from below:")
	fmt.Println("+ - * /")

    _,err3 := fmt.Scanf("%s", &operation)
    CheckErrorNil(err3)
	operation1 := strings.TrimSpace(operation)
if len(operation1)==1{
	if operation1=="+"{
		result:=Add(val1,val2)
		fmt.Println(result)
	}else if operation1=="-"{
		result:=Subtract(val1,val2)
		fmt.Println(result)

	}else if operation1=="*"{
		result:=Multiply(val1,val2)
		fmt.Println(result)

	}else if operation1=="/"{
        result:=Divide(val1,val2)
		fmt.Println(result)

	}else{
		fmt.Println("Invalid Operation")
	}
}else{
	fmt.Println("Only one Operation is acceptable")
}
}
func CheckErrorNil(err error){
	if err!=nil{
    fmt.Println(err)
	}
}
func Add(a,b float64)float64{
return a + b
}
func Subtract(a,b float64)float64{
return a-b
}
func Multiply(a,b float64)float64{
	return a*b
}
func Divide(a,b float64)float64{
	return a/b
}