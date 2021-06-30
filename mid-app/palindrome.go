package main

import "fmt"

func isPalindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}

	}

	return true
}

func main() {
	var input string

	fmt.Println("Please input the word: ")
	fmt.Scanf("%s\n", &input)
	if input != "" {
		var check bool = isPalindrome(input)
		if check == true {
			fmt.Println(input, "is palindrome")
		} else {
			fmt.Println(input, "is not palindrome")
		}
	} else {
		fmt.Println("-------------")
		fmt.Println("Please input the word!")
	}
}
