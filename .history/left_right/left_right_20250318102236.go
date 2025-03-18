package left_right

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decodeSequence(encodedStr string) string {
	startNum := 0
	if encodedStr[0] == 'L' {
		startNum = 2
	}

	number := []int{startNum}

	for i := 0; i < len(encodedStr); i++ {
		switch encodedStr[i] {
		case 'L':
			number = append(number, number[i]-1)
		case 'R':
			number = append(number, number[i]+1)
		case '=':
			number = append(number, number[i])
		}
	}

	var result strings.Builder
	for _, num := range number {
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