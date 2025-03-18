package main

import (
	"fmt"
)

func main() {
	var pattern string
	fmt.Print("Enter the pattern: ")
	fmt.Scanln(&pattern)

	result := DecodePattern(pattern)
	fmt.Println("Best number with minimum sum:", result)
}