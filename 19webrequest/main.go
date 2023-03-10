package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url = "https://app.open.money"
func main() {
	fmt.Println("Welcome to WebRequest in Go-lang  : ")
response,err:=http.Get(url)
checkErrorNil(err)
fmt.Printf("Response is of Type : %T\n",response)
defer response.Body.Close()
databyte,err:=ioutil.ReadAll(response.Body)
checkErrorNil(err)
fmt.Println("Response Body is :",string(databyte))
}
func checkErrorNil(err error){
	if err !=nil{
		panic(err)
	}
}