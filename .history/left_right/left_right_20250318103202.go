package left_right

import (
	"math"
	"strconv"
)



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
func DecodePattern(pattern string) string {
	minSumObj := struct {
		minSum   int
		bestNumber string
	}{minSum: math.MaxInt, bestNumber: ""}
	for digit := 0; digit <= 9; digit++ {
		generateNumber(pattern, 1, strconv.Itoa(digit), &minSumObj)
	}
	return minSumObj.bestNumber
}