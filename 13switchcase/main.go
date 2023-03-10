package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Welcome to switchcase class:")
	rand.Seed(time.Now().Unix())
	diceNumber:=rand.Intn(6) +1

fmt.Println("value of diceNumber is :",diceNumber)

switch diceNumber{
case 1:
	fmt.Println("diceNumber is 1")
case 2:
	fmt.Println("diceNumber is 2")
case 3:
	fmt.Println("diceNumber is 3")
case 4:
	fmt.Println("diceNumber is 4")
case 5:
	fmt.Println("diceNumber is 5")
case 6:
	fmt.Println("diceNumber is 6 you can roll again")
default:
	fmt.Println("What was that")
	
}
}