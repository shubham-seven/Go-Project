package main

import (
	"encoding/json"
	"fmt"
	//"os"
	//"strconv"
)

type ColourGroup struct {
	ID    int
	Name  string
	Color string
}

func main() {
/*
	group1 := ColourGroup{
		ID:    1,
		Name:  "shubham",
		Color: "Red",
	}
	group2 := ColourGroup{
		ID:    2,
		Name:  "sumit",
		Color: "black",
	}
	p, _ := strconv.ParseFloat("123.9079087", 64)
	fmt.Println(p)
	fmt.Println(group1)
	fmt.Println(group2)

	x,err := json.Marshal(group1)
	if err!=nil{
		fmt.Println("error:",err)
	}
	os.Stdout.Write(x)
	*/

//Unmarshall json 
var jsonblob = []byte(`[
{"ID":    1,"Name":  "shubham","Color": "Red"},
{"ID":    2,"Name":  "sumit","Color": "black"}
]`)
var cgr []ColourGroup
err1:=json.Unmarshal(jsonblob,&cgr)
if err1!=nil{
	fmt.Println(err1)
}
fmt.Printf("%+v",cgr)
}
