package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads a file containing lines of integers and returns a 2D array of integers
// File looking like:
//
//	---
//	 1 2 3
//	 4 5 6
//	 ---
//
// Gets turnes into [[1,2,3],[4,5,6]]
func Read2dIntArray(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var currentLine []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				panic("Can not convert string to integer")
			}
			currentLine = append(currentLine, num)
		}

		// Ignore empty lines, just to be sure
		if len(currentLine) > 0 {
			result = append(result, currentLine)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return result, nil
}

func ReadString(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}

	return string(file), nil
}
