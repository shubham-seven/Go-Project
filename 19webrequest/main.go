package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
)
const url = "http://services.explorecalifornia.org/json/tours.php"
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
func toursFromJson(content string)[] Tour{
tours:=make([]Tour,0,20)
decoder:=json.NewDecoder(strings.NewReader(content))
_,err:=decoder.Token()
checkErrorNil(err)
}
var tour Tour 
for decoder.More(){
	err:=decoder.Decode(&tour)
	checkErrorNil(err)
tours:=append(tours,tour)
}
type Tour struct{
	Name,Price string
}
