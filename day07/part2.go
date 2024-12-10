package day07

import (
	"fmt"
	"strconv"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

// Returns a new number if we string append addition to the original
// appendValue(123, 45) -> 12345
func appendValue(original int, addition int) int {
	val, _ := strconv.Atoi(strconv.Itoa(original) + strconv.Itoa(addition))
	return val
}

// DFS through solution space
func solveEquationTri(currentValue int, operator string, nums []int, result int) bool {
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
	} else if operator == "||" {
		currentValue = appendValue(currentValue, head)
	}

	// if we are already over result number, we can safely prune solution tree here
	if currentValue > result {
		return false
	}

	matchAdd := solveEquationTri(currentValue, "+", tail, result)
	matchMul := solveEquationTri(currentValue, "*", tail, result)
	matchAppend := solveEquationTri(currentValue, "||", tail, result)

	return matchAdd || matchMul || matchAppend
}

func Part2() {
	equations, _ := utils.Read2dIntArray("day07/test.txt")

	sum := 0
	for _, eq := range equations {
		result := eq[0]
		nums := eq[1:]

		if solveEquationTri(0, "+", nums, result) {
			sum += result
			fmt.Printf("Solved %d = %v\n", result, nums)
		}
	}
	fmt.Println("Total sum of solvable results", sum)
}
