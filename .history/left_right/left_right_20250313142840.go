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
	
	// Initialize first number based on first character
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

		switch curr {
		case 'L':
			if prev == '=' && numbers[i] == 0 {
				numbers[i+1] = 2
			} else if prev == 'R' {
				numbers[i+1] = numbers[i] - 1
			} else {
				numbers[i+1] = numbers[i] - 1
			}
		case 'R':
			if prev == '=' && next == 'L' {
				numbers[i+1] = 0
			} else {
				numbers[i+1] = numbers[i] + 1
			}
		case '=':
			if next == 'L' {
				numbers[i+1] = numbers[i]
			} else if prev == '=' {
				numbers[i+1] = numbers[i]
			} else {
				numbers[i+1] = numbers[i]
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