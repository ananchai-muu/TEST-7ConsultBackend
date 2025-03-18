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
	
	// Initialize first number and handle first character
	switch encodedStr[0] {
	case 'L':
		numbers[0] = 2
		numbers[1] = 1
	case 'R':
		numbers[0] = 0
		numbers[1] = 1
	case '=':
		numbers[0] = 0
		numbers[1] = 0
	}

	// Process remaining characters
	for i := 1; i < len(encodedStr); i++ {
		curr := encodedStr[i]
		prev := encodedStr[i-1]
		
		// Get next character if available
		var next byte = '='
		if i < len(encodedStr)-1 {
			next = encodedStr[i+1]
		}

		// Get previous value
		prevValue := numbers[i]

		// Calculate next value based on patterns
		switch curr {
		case 'L':
			if prev == '=' && prevValue == 0 {
				// If previous is = and current value is 0, start from 2
				numbers[i+1] = 2
			} else if prev == 'R' {
				// If coming from R, decrease by 1
				numbers[i+1] = prevValue - 1
			} else {
				// Normal L case, decrease by 1 but not below 0
				nextValue := prevValue - 1
				if nextValue < 0 {
					nextValue = 0
				}
				numbers[i+1] = nextValue
			}
		case 'R':
			if prev == '=' && next == 'L' {
				// If between = and L, stay at 0
				numbers[i+1] = 0
			} else if prev == 'L' {
				// If coming from L, start increasing from current value
				numbers[i+1] = prevValue + 1
			} else {
				// Normal R case, increase by 1
				numbers[i+1] = prevValue + 1
			}
		case '=':
			if prev == 'R' && next == 'R' {
				// If between R and R, maintain the sequence
				numbers[i+1] = prevValue
			} else {
				// Normal = case, keep previous value
				numbers[i+1] = prevValue
			}
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