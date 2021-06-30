package main

import "fmt"

func main() {
	// var name1 = "eko"
	// var name2 = "eko"

	// var result bool = name1 == name2
	// var result2 bool = name1 != name2

	// fmt.Println(result, result2)

	//array
	// var arrNames [3]string
	// arrNames[0] = "herza"
	// arrNames[1] = "parama"
	// arrNames[2] = "yudhanto"

	// fmt.Println(arrNames)

	var values = [5]int{
		12,
		13,
		14,
		15,
		16,
	}
	//len untuk panjang array bukan jumlah data
	fmt.Println(len(values))

}
