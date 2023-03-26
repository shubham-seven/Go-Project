package main

import "fmt"

func main(){
	count:=-1
i := 1992
for i<=2023{
	count++
	i++
}
fmt.Printf("your age: %v\t\n",count)
j:=1992
flag:=0
for{
if j>=2023{
	break
}
flag++
j++
}
fmt.Printf("your age : %v\n",flag)
}