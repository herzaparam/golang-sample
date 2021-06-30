package main

import "fmt"

func main() {
	// var months = [...]string{
	// 	"januari",
	// 	"februari",
	// 	"maret",
	// 	"april",
	// 	"mei",
	// 	"juni",
	// 	"juli",
	// 	"agustus",
	// 	"september",
	// 	"oktober",
	// 	"november",
	// 	"desember",
	// }

	// slice1 := months[10:]
	// slice1[0] = "mei-"
	// slice1[1] = "juni-"
	// fmt.Println(months)

	// slice2 := append(slice1, "bulan13")
	// slice2[0] = "haha"
	// fmt.Println(slice2)
	// fmt.Println(months)

	slice3 := make([]string, 3, 5)
	slice3[0] = "her"
	slice3[1] = "za"
	slice3[2] = "par"

	// fmt.Println(slice3)

	copySlice := make([]string, len(slice3), cap(slice3))
	copy(copySlice, slice3)
	fmt.Println(copySlice)
}
