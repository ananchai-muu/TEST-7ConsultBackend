package left_right

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decodeSequence(encodedStr string) string {
	numbers := make([]int, len(encodedStr)+1)
	
	// Count consecutive characters at the start
	equalCount := 0
	for i := 0; i < len(encodedStr) && encodedStr[i] == '='; i++ {
		equalCount++
	}

	// Initialize based on pattern
	if equalCount > 0 {
		// If starts with =, initialize all positions up to equalCount as 0
		for i := 0; i <= equalCount; i++ {
			numbers[i] = 0
		}
	} else {
		// Initialize based on first character
		switch encodedStr[0] {
		case 'L':
			numbers[0] = 2
			numbers[1] = 1
		case 'R':
			numbers[0] = 0
			numbers[1] = 1
		}
	}

	// Process remaining characters
	for i := equalCount; i < len(encodedStr); i++ {
		curr := encodedStr[i]
		
		// Get previous character and value
		prev := byte('=')
		if i > 0 {
			prev = encodedStr[i-1]
		}
		prevValue := numbers[i]

		// Get next character if available
		next := byte('=')
		if i < len(encodedStr)-1 {
			next = encodedStr[i+1]
		}

		switch curr {
		case 'L':
			if prev == '=' {
				// After =, start decreasing from 2
				numbers[i+1] = 2
			} else if prev == 'R' {
				// After R, decrease from current value
				numbers[i+1] = prevValue - 1
			} else {
				// Normal case, decrease but not below 0
				nextValue := prevValue - 1
				if nextValue < 0 {
					nextValue = 0
				}
				numbers[i+1] = nextValue
			}
		case 'R':
			if prev == '=' && next == 'L' {
				// Between = and L, use 0
				numbers[i+1] = 0
			} else if prev == 'L' {
				// After L, start from 0
				numbers[i+1] = 0
			} else if next == 'L' {
				// Before L, use 2
				numbers[i+1] = 2
			} else {
				// Normal case, increase
				numbers[i+1] = prevValue + 1
			}
		case '=':
			// Keep previous value
			numbers[i+1] = prevValue
		}
	}

	// Convert to string
	var result strings.Builder
	for _, num := range numbers {
		result.WriteString(strconv.Itoa(num))
	}
	return result.String()
}

func sumOfDigits(numberStr string) int {
	sum := 0
	for _, char := range numberStr {
		num, _ := strconv.Atoi(string(char))
		sum += num
	}
	return sum
}

func FindBestNumberFromEncoded(encodedStr string) (string, int) {
	decodedNumber := decodeSequence(encodedStr)
	result := sumOfDigits(decodedNumber)
	return decodedNumber, result
}

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string (L/R/=): ")
	encodedStr, _ := reader.ReadString('\n')
	return strings.TrimSpace(encodedStr)
}