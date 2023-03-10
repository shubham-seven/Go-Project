package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func main() {
	fmt.Println("Welcome to Files in Go-lang  : ")

content := "The name of the File is learncodeonline.in"

file,err := os.Create("./mylcofile.txt")
checkNilErr(err)
length,err:=io.WriteString(file,content)
checkNilErr(err)
fmt.Println("Length of the file is : ",length)
defer file.Close()
readFile("./mylcofile.txt")
}
func readFile(filename string){
databyte,err:=ioutil.ReadFile(filename)
checkNilErr(err)
fmt.Println("Text Data inside the file is :",databyte)
fmt.Println("Text Data inside the file is :",string(databyte))

}
func checkNilErr(err error){
if err !=nil{
	panic(err)
}
}