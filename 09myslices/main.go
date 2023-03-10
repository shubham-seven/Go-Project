package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Welcome to the class of Slices")
	

  var fruitlist = []string{"mango","water-melon","banana","apple"}
  fmt.Printf("Type of fruitlist is %T\n",fruitlist)

  fmt.Println(fruitlist)

  fruitlist = append(fruitlist,"pineapple","peach")

fmt.Println(fruitlist)

fruitlist = append(fruitlist[1:3])

fmt.Println(fruitlist)

  veglist := make([]int,4)

  veglist[0]=111
  veglist[1]=999
  veglist[2]=500
  veglist[3]=100

  fmt.Println(veglist)

  sort.Ints(veglist)

  fmt.Println(veglist)

  //How to remove a value from slices based on index

  var courses = []string{"reactjs","javascript","python","Ruby","Go-lang"}

  fmt.Println(courses)
  var index int = 2
  
  courses = append(courses[:index],courses[index+1:]...)

fmt.Println(courses)
  
}