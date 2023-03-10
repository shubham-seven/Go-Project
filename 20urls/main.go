package main

import (
	"fmt"
	"net/url"
	
)

const myUrl string="https://lco.dev:3000/learn?coursename=reactjs&paymentid=asdf234jkl"

func main(){
	fmt.Println("Welcome to the class of URLs :")
fmt.Println(myUrl)
//parsing
result,err:=url.Parse(myUrl)
checkNilError(err)
fmt.Println(result.Scheme)
fmt.Println(result.RawQuery)
fmt.Println(result.Port())
fmt.Println(result.Host)
fmt.Println(result.Path)

qparams:=result.Query()
fmt.Printf("The type of the query params are: %T\n",qparams)
fmt.Println(qparams["coursename"])
fmt.Println(qparams["paymentid"])

for _,val:=range qparams{
	fmt.Println("The params is :",val)
}
partsofurl := &url.URL{
	Scheme: "https",
	Host: "lco.dev",
	Path: "/tutcss",
	RawQuery: "user=hitesh",
}
anotherUrl:=partsofurl.String()

fmt.Println(anotherUrl)

}
func checkNilError(err error){
	if err !=nil{
		panic(err)
	}
}