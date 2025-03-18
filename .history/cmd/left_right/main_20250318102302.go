package main

import (
	"fmt"
	"strings"

	"backendtest/left_right"
)

func main() {
	encodedStr := left_right.GetUserInput()
	if len(encodedStr) == 0 || strings.ContainsAny(encodedStr, "0123456789") {
		fmt.Println("Invalid input! Please enter only 'L', 'R', or '='.")
		return
	}
	
	decodedNumber, sum := left_right.FindBestNumberFromEncoded(encodedStr)
	fmt.Println("\nEncoded:", encodedStr)
	fmt.Println("Decoded Number:", decodedNumber)
	fmt.Println("Sum of Digits:", sum)
} 