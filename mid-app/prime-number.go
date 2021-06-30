package main

import (
	"fmt"
	"math"
)

func isPrime(num1 int, num2 int) {
	if num1 < 1 || num2 < 1 {
		fmt.Println("invalid number of input")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			if num1 != 1 {
				fmt.Printf("%d ", num1)
			}
		}
		num1++
	}
	fmt.Println()
}

func main() {
	var numInput1 int
	var numInput2 int

	fmt.Println("Please input the first number: ")
	fmt.Scanf("%d\n", &numInput1)
	fmt.Println("Please input the second number: ")
	fmt.Scanf("%d\n", &numInput2)
	fmt.Println("prime number between", numInput1, "to", numInput2, "are :")
	isPrime(numInput1, numInput2)

}
