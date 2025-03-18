package main

import (
	"backendtest/left_right"
	"fmt"
)

func main() {
	var pattern string
	fmt.Print("Enter the pattern: ")
	fmt.Scanln(&pattern)

	result := left_right.DecodePattern(pattern)
	fmt.Println("Best number with minimum sum:", result)
}