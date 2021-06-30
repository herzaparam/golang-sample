package main

import (
	"fmt"
)

func main() {

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println(counter)
	// }

	//for ranges iterasi untuk array, slice dan map
	// names := []string{"herza", "parama", "yudhanto"}
	// for _, name := range names {
	// 	fmt.Println(name)
	// }
	// fmt.Println(reflect.TypeOf(names))

	person := make(map[string]string)
	person["name"] = "Herza"
	person["age"] = "25"
	person["job"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
