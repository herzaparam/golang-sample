package main

import "fmt"

func main() {
	var number1 int8
	var number2 int8
	var operator string

	fmt.Println("enter your first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("enter your second number")
	fmt.Scanf("\n%d", &number2)
	fmt.Println("enter the operator")
	fmt.Scanf("\n%s", &operator)

	switch operator {
	case "+":
		plus(number1, number2)
	case "-":
		minus(number1, number2)
	case "/":
		divider(number1, number2)
	case "*":
		multiplier(number1, number2)
	default:
		fmt.Println("please insert your operator")
	}

}

func plus(num1 int8, num2 int8) {
	num := num1 + num2
	fmt.Println(num)
}

func minus(num1 int8, num2 int8) {
	num := num1 - num2
	fmt.Println(num)
}

func divider(num1 int8, num2 int8) {
	num := num1 / num2
	fmt.Println(num)
}

func multiplier(num1 int8, num2 int8) {
	num := num1 * num2
	fmt.Println(num)
}
