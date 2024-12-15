package day10

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func TrailheadScore(i, j, current int, reached *mapset.Set[utils.Location], grid *[][]int) int {

	g := *grid
	if current == 9 {
		(*reached).Add(utils.Location{I: i, J: j})
		return 1
	}

	next := current + 1

	// up
	if i >= 1 && g[i-1][j] == next {
		TrailheadScore(i-1, j, next, reached, grid)
	}

	// right
	if j < len(g[i])-1 && g[i][j+1] == next {
		TrailheadScore(i, j+1, next, reached, grid)
	}

	// down
	if i < len(g)-1 && g[i+1][j] == next {
		TrailheadScore(i+1, j, next, reached, grid)
	}

	// left
	if j >= 1 && g[i][j-1] == next {
		TrailheadScore(i, j-1, next, reached, grid)
	}

	return (*reached).Cardinality()
}

func Part1() {

	input, _ := utils.ReadString("day10/test.txt")

	grid, err := utils.Sto2dIntArray(input, "")
	if err != nil {
		fmt.Println("Error reading input as 2dintArray", err)
	}
	utils.Print2dArray(grid)

	scoreSum := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				reached := mapset.NewSet[utils.Location]()
				score := TrailheadScore(i, j, 0, &reached, &grid)

				fmt.Printf("Trailhead at (%d,%d) score = %d\n", i, j, score)
				fmt.Println("-----------")
				scoreSum += score
			}
		}
	}

	fmt.Println("Total Score", scoreSum)
}
