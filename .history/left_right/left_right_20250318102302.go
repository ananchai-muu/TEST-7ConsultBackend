package left_right

import (
	"bufio"
	"fmt"
	"math"
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

func getSum(number string) int {
	sum := 0
	for _, digit := range number {
		d, _ := strconv.Atoi(string(digit))
		sum += d
	}
	return sum
}

// Function to check if a transition is valid based on the condition
func isValidTransition(prevDigit, nextDigit int, condition string) bool {
	if condition == "L" && prevDigit > nextDigit {
		return true
	} else if condition == "R" && prevDigit < nextDigit {
		return true
	} else if condition == "=" && prevDigit == nextDigit {
		return true
	}
	return false
}

// Recursive function to generate a number based on the pattern
func generateNumber(pattern string, index int, current string, minSumObj *struct {
	minSum   int
	bestNumber string
}) {
	if index == len(pattern)+1 {
		sum := getSum(current)
		if sum < minSumObj.minSum {
			minSumObj.minSum = sum
			minSumObj.bestNumber = current
		}
		return
	}

	for digit := 0; digit <= 9; digit++ {
		if index == 0 || isValidTransition(int(current[len(current)-1]-'0'), digit, string(pattern[index-1])) {
			generateNumber(pattern, index+1, current+strconv.Itoa(digit), minSumObj)
		}
	}
}

// Function to decode the pattern
func decodePattern(pattern string) string {
	minSumObj := struct {
		minSum   int
		bestNumber string
	}{minSum: math.MaxInt, bestNumber: ""}
	for digit := 0; digit <= 9; digit++ {
		generateNumber(pattern, 1, strconv.Itoa(digit), &minSumObj)
	}
	return minSumObj.bestNumber
}