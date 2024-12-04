package day04

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func isMatch(word string) bool {
	return word == "MAS" || word == "SAM"
}

func Part2() {
	puzzle, err := utils.Read2dCharArray("day04/input.txt")
	if err != nil {
		fmt.Println("Error opening file as 2darray")
		return
	}

	foundCount := 0
	for i := 1; i < len(puzzle)-1; i++ {
		for j := 1; j < len(puzzle[i])-1; j++ {

			center := puzzle[i][j]

			first := puzzle[i-1][j-1] + center + puzzle[i+1][j+1]
			second := puzzle[i-1][j+1] + center + puzzle[i+1][j-1]

			if isMatch(first) && isMatch(second) {
				foundCount += 1
			}

		}
	}

	fmt.Println("Total hits count", foundCount)

}
