package main

import "fmt"

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "shubham",
		friends:   map[string]int{"vivek": 90, "rajat": 80, "chiku": 70},
		favDrinks: []string{"milkShake", "Water"},
	}
	fmt.Println(s)

	for k, v := range s.friends {
		fmt.Printf("Friends name : %s\t contactNo : %v\n", k, v)
	}
	for j, val := range s.favDrinks {
		fmt.Printf("Index is : %v\t value is : %s\n", j, val)
	}
}
