package day10

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func TrailheadRating(i, j, current int, grid [][]int) int {

	if current == 9 {
		return 1
	}

	next := current + 1
	rating := 0

	// up
	if i >= 1 && grid[i-1][j] == next {
		rating += TrailheadRating(i-1, j, next, grid)
	}

	// right
	if j < len(grid[i])-1 && grid[i][j+1] == next {
		rating += TrailheadRating(i, j+1, next, grid)
	}

	// down
	if i < len(grid)-1 && grid[i+1][j] == next {
		rating += TrailheadRating(i+1, j, next, grid)
	}

	// left
	if j >= 1 && grid[i][j-1] == next {
		rating += TrailheadRating(i, j-1, next, grid)
	}

	return rating
}

func Part2() {

	input, _ := utils.ReadString("day10/input.txt")

	grid, err := utils.Sto2dIntArray(input, "")
	if err != nil {
		fmt.Println("Error reading input as 2dintArray", err)
	}
	utils.Print2dArray(grid)

	ratingSum := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				rating := TrailheadRating(i, j, 0, grid)

				fmt.Printf("Trailhead at (%d,%d) score = %d\n", i, j, rating)
				fmt.Println("-----------")
				ratingSum += rating
			}
		}
	}

	fmt.Println("Total Rating", ratingSum)
}
