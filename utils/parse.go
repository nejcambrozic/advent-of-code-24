package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Sto2dIntArray(input string, delimiter string) ([][]int, error) {
	result := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		elements := strings.Split(line, delimiter)

		currentElements := []int{}
		for _, element := range elements {
			el, err := strconv.Atoi(element)
			if err != nil {
				fmt.Println("Error parsing int", err)
				//panic("Failed to parse int")
				el = -1
			}
			currentElements = append(currentElements, el)
		}

		if len(currentElements) > 0 {

			result = append(result, currentElements)
		}

	}
	return result, nil
}
