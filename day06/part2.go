package day06

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func (p1 *Position) Equals(p2 *Position) bool {
	return p1.Row == p2.Row && p1.Column == p2.Column && p1.Orientation == p2.Orientation
}

func detectLoop(grid [][]string, pos Position) bool {
	path := []Position{}
	var err error
	for true {
		pos, err = step(grid, pos)

		if err != nil {

			return false
		}

		for _, visited := range path {
			if pos == visited {
				return true
			}
		}
		path = append(path, pos)

	}
	return false

}

func Part2() {

	grid, err := utils.Read2dCharArray("day06/input.txt")
	if err != nil {
		fmt.Println("Error reading data as 2d char array")
		return
	}
	startingPosition, err := getCurrentPosition(grid)
	if err != nil {
		return
	}

	path, _ := walkTheGrid(grid, startingPosition)
	fmt.Println("visitedNum", len(path))

	loopCount := 0
	startingPos := path[0]
	for _, position := range path {
		if position.Row == startingPos.Row && position.Column == startingPos.Column {
			continue
		}

		iterationGrid := utils.DeepCopy2dStrSpice(grid)
		iterationGrid[position.Row][position.Column] = "#" // place an obstacle on path traveled

		if detectLoop(iterationGrid, startingPos) {
			loopCount += 1
		}

	}
	fmt.Println("Could produce loops", loopCount)

}
