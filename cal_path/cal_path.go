package cal_path

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data [][]int

func CalMaxPath(triangle Data) int {
	if len(triangle) == 0 {
		return 0
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += max(triangle[row+1][col], triangle[row+1][col+1])
		}
	}

	return triangle[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ReadJSON(filename string) (Data, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data Data
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return data, nil
}

func main() {
	data, err := ReadJSON("cal_path/hard.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result := CalMaxPath(data)

	fmt.Println("Maximum Path Sum:", result)
}
