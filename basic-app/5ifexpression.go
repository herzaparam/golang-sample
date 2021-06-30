package main

import "fmt"

func main() {
	name := "haha"

	// if name == "Herza" {
	// 	fmt.Println(name)
	// } else if name == "herza" {
	// 	fmt.Println("name salah")
	// } else {
	// 	fmt.Println("haha")
	// }

	// if length := len(name); length >= 4 {
	// 	fmt.Println("berhasil")
	// } else {
	// 	fmt.Println("gagal")
	// }

	//switch
	switch name {
	case "haha":
		fmt.Println("name = haha")
	case "HaHa":
		fmt.Println("name = HaHa")
	}

}
