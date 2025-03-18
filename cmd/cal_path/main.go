package main

import (
	"fmt"
	"log"

	"backendtest/cal_path"
)

func main() {
	data, err := cal_path.ReadJSON("cal_path/hard.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result := cal_path.CalMaxPath(data)
	fmt.Println("Maximum Path Sum:", result)
} 