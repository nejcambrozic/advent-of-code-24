package day12

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func Part2() {
	inputMap, _ := utils.Read2dCharArray("day12/input.txt")

	gridSize := len(inputMap)

	queue := Queue{}

	visited := make([][]bool, gridSize)
	for i := range visited {
		visited[i] = make([]bool, gridSize)
	}
	totalPrice := 0

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	diagonalDirections := map[int][]int{
		0: {-1, 1},
		1: {1, 1},
		2: {1, -1},
		3: {-1, -1},
	}

	for i := range inputMap {
		for j := range inputMap[i] {

			if visited[i][j] {
				continue
			}
			visited[i][j] = true

			loc := utils.Location{I: i, J: j}
			currPlot := inputMap[i][j]

			queue.Enqueue(loc)
			sides := 0
			area := 1

			for !queue.IsEmpty() {
				loc, _ := queue.Dequeue()

				// Misses for each direction: up, right, down, left
				// Adding latest 5th element that get's the value of first so we can compare in pairs
				misses := [5]bool{false, false, false, false, false}
				for z, dr := range directions {
					nextL := utils.Location{I: loc.I + dr[0], J: loc.J + dr[1]}

					if utils.IsValidLocation(nextL, gridSize) && inputMap[nextL.I][nextL.J] == currPlot {
						if !visited[nextL.I][nextL.J] {
							queue.Enqueue(nextL)
							visited[nextL.I][nextL.J] = true
							area += 1
						}
					} else {
						misses[z] = true
					}
				}
				misses[4] = misses[0]

				for m := 0; m < 4; m++ {
					// Outercorner: 2 consequative misses
					if misses[m] && misses[m+1] {
						sides += 1
					}
					// Innercorner: when 2 consequative hits, but the element in between doesn't match the plot
					if !misses[m] && !misses[m+1] {
						locMod, _ := diagonalDirections[m]

						dLoc := utils.Location{I: loc.I + locMod[0], J: loc.J + locMod[1]}
						if utils.IsValidLocation(dLoc, gridSize) && inputMap[dLoc.I][dLoc.J] != currPlot {
							sides += 1
						}
					}
				}
			}
			fencePrice := area * sides
			fmt.Printf("Plot %s (%d * %d) = %d\n", currPlot, area, sides, fencePrice)
			totalPrice += fencePrice
		}
	}

	fmt.Println("Total price", totalPrice)
}
