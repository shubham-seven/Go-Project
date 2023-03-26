package main

import "fmt"

func main() {

	m := map[string][]string{
		"tiwari":  []string{"harshit", "dubey", "manish"},
		"shubham": []string{"sumit", "talreja", "trivedi"},
	}

	for k, v := range m {
		fmt.Printf("key is : %s\n", k)

		for i, val := range v {
			fmt.Printf("index is : %v\t and value is : %s\n", i, val)
		}
	}
	m["fulo"] = []string{"tobbaco", "liqour", "cigars"}

	for k, v := range m {
		fmt.Printf("key is : %s\n", k)

		for i, val := range v {
			fmt.Printf("index is : %v\t and value is : %s\n", i, val)
		}
	}
	delete(m, "fulo")
	fmt.Println(m)

}
