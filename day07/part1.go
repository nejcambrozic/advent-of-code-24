package day07

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

// DFS through solution space
func solveEquationBinary(currentValue int, operator string, nums []int, result int) bool {

	// reached the end of tree: see if we got the correct result
	if len(nums) == 0 {
		return currentValue == result
	}

	head := nums[0]
	tail := nums[1:]
	if operator == "+" {
		currentValue += head
	} else if operator == "*" {
		currentValue *= head
	}

	// if we are already over result number, we can safely prune solution tree here
	if currentValue > result {
		return false
	}

	matchAdd := solveEquationBinary(currentValue, "+", tail, result)
	matchMul := solveEquationBinary(currentValue, "*", tail, result)

	return matchAdd || matchMul
}

func Part1() {
	equations, _ := utils.Read2dIntArray("day07/test.txt")

	sum := 0
	for _, eq := range equations {
		result := eq[0]
		nums := eq[1:]

		if solveEquationBinary(0, "+", nums, result) {
			sum += result
			fmt.Printf("Solved %d = %v\n", result, nums)
		}
	}
	fmt.Println("Total sum of solvable results", sum)
}
