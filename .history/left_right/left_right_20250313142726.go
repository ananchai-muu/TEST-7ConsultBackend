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
	
	// กำหนดค่าเริ่มต้น
	if encodedStr[0] == 'L' {
		numbers[0] = 2
	} else {
		numbers[0] = 0
	}

	// คำนวณค่าที่เหลือ
	for i := 0; i < len(encodedStr); i++ {
		// ดูค่าก่อนหน้า (ถ้ามี)
		prevValue := 0
		if i > 0 {
			prevValue = numbers[i]
		}

		// ดูตัวอักษรถัดไป (ถ้ามี)
		nextChar := byte('=')
		if i < len(encodedStr)-1 {
			nextChar = encodedStr[i+1]
		}

		// คำนวณค่าตามเงื่อนไข
		switch encodedStr[i] {
		case 'L':
			if i == 0 {
				// ถ้าเป็นตัวแรกและเป็น L
				numbers[i+1] = 1
			} else if prevValue == 0 && nextChar == 'L' {
				// ถ้าค่าก่อนหน้าเป็น 0 และตัวถัดไปเป็น L
				numbers[i+1] = 2
			} else {
				// กรณีอื่นๆ ลดค่าลง 1
				numbers[i+1] = prevValue - 1
			}
		case 'R':
			if i == 0 {
				// ถ้าเป็นตัวแรกและเป็น R
				numbers[i+1] = 1
			} else if prevValue == 0 && nextChar == 'L' {
				// ถ้าค่าก่อนหน้าเป็น 0 และตัวถัดไปเป็น L
				numbers[i+1] = 0
			} else {
				// กรณีอื่นๆ เพิ่มค่าขึ้น 1
				numbers[i+1] = prevValue + 1
			}
		case '=':
			if i == 0 {
				// ถ้าเป็นตัวแรกและเป็น =
				numbers[i+1] = 0
			} else if nextChar == 'L' && prevValue == 0 {
				// ถ้าตัวถัดไปเป็น L และค่าก่อนหน้าเป็น 0
				numbers[i+1] = 0
			} else {
				// กรณีอื่นๆ คงค่าเดิม
				numbers[i+1] = prevValue
			}
		}
	}

	// แปลงเป็น string
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