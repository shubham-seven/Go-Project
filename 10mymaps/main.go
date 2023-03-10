package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of maps : ")

	languages := make(map[string]string)

	languages["JS"] = "ja"
	languages["RB"] = "R"
	languages["PY"] = "P"

	fmt.Println(languages)

	//delete(languages,"JS")
	fmt.Println(languages)

	//Loops are interesting in go-lang

	for key, value := range languages {
		fmt.Printf("For key  %v , value is %v\n", key, value)
	}

}
