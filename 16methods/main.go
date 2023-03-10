package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Methods  : ")
//there is no inheritance,no super or parent or child class in GO-LANG
structure := User{"shubham","shubhamrolly@gmail.com",true,31}

//fmt.Println(structure)

fmt.Printf("All Details of User are : %+v\n",structure)
fmt.Printf("Name is : %v and E-mail is : %v and Age is : %v\n",structure.Name,structure.Email,structure.age)
structure.GetStatus()
structure.NewEmail()
fmt.Printf("E-mail is : %v\n",structure.Email)

}

type User struct{
Name string
Email string
status bool
age int
}
func(u User) GetStatus(){
fmt.Println("Is User Active : ",u.status)
}
func(u User) NewEmail(){
	u.Email = "XYZ@gmail.com"
	fmt.Println("User Email is : ",u.Email)
}
