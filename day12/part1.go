package day12

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

type Queue []utils.Location

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(l utils.Location) {
	*q = append(*q, l)
}

func (q *Queue) Dequeue() (utils.Location, bool) {
	if q.IsEmpty() {
		return utils.Location{I: -1, J: -1}, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func Part1() {
	inputMap, _ := utils.Read2dCharArray("day12/input.txt")

	gridSize := len(inputMap)

	queue := Queue{}

	visited := make([][]bool, gridSize)
	for i := range visited {
		visited[i] = make([]bool, gridSize)
	}
	totalPrice := 0

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := range inputMap {
		for j := range inputMap[i] {

			if visited[i][j] {
				continue
			}
			visited[i][j] = true

			loc := utils.Location{I: i, J: j}
			currPlot := inputMap[i][j]

			queue.Enqueue(loc)
			perimeter := 0
			area := 1
			fmt.Println("Starting new plot", currPlot, loc)

			for !queue.IsEmpty() {
				loc, _ := queue.Dequeue()

				for _, dr := range directions {
					nextL := utils.Location{I: loc.I + dr[0], J: loc.J + dr[1]}

					if utils.IsValidLocation(nextL, gridSize) && inputMap[nextL.I][nextL.J] == currPlot {
						if !visited[nextL.I][nextL.J] {
							queue.Enqueue(nextL)
							visited[nextL.I][nextL.J] = true
							area += 1
						}
					} else {
						perimeter += 1
					}
				}
			}
			fencePrice := area * perimeter
			fmt.Printf("Plot %s (%d * %d) = %d\n", currPlot, area, perimeter, fencePrice)
			totalPrice += fencePrice
		}
	}

	fmt.Println("Total price", totalPrice)
}
