package main


import "fmt"


func main(){

	slice1 := [] string{"shubham","sumit","talreja","trivedi"}
	slice2 := [] string{"tiwari","harshit","dubey","manish"}
	slice3 := [] [] string{slice1,slice2}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	for i,slice4:=range slice3{
		fmt.Printf("record : %v\n",i)
		for j,val:=range slice4{
			fmt.Printf("Index is : %v\t value is : %v\n",j,val)
		}
	}


}