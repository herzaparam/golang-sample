package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Herza",
		"age":  "25",
	}

	person["address"] = "wage"

	fmt.Println(len(person))
}
