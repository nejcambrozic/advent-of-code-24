package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func MulSum(instruction string) int {
	mulSum := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(instruction, -1)
	for _, match := range matches {
		// match will be 3 el list containing
		// [entireMatchedString, firstNumAsString, secondNumAsString]
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		mulSum += first * second
	}
	return mulSum
}

func Part1() {
	instructions, err := utils.ReadString("day03/input.txt")
	if err != nil {
		fmt.Println("Error reading input as single string", err)
		return
	}

	mulSum := MulSum(instructions)
	fmt.Println("Total sum: ", mulSum)
}
