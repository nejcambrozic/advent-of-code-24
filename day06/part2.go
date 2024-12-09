package day06

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func (p1 *Position) Equals(p2 *Position) bool {
	return p1.Row == p2.Row && p1.Column == p2.Column && p1.Orientation == p2.Orientation
}

func detectLoop(grid [][]string, pos Position) bool {
	path := mapset.NewSet[Position]()
	var err error
	for true {
		pos, err = step(grid, pos)

		// not a loop if we manage to exit the grid
		if err != nil {
			return false
		}
		// if we come to same position in same orientation: we are in a loop
		visited := path.Contains(pos)
		if visited {
			return true
		}
		path.Add(pos)
	}
	return false
}

func Part2() {

	grid, err := utils.Read2dCharArray("day06/input.txt")
	if err != nil {
		fmt.Println("Error reading data as 2d char array")
		return
	}

	// first walk the grid so we get the path walked in part1
	// we will try to produce a path by changing each
	startingPosition, err := getCurrentPosition(grid)
	if err != nil {
		return
	}
	path, _ := walkTheGrid(grid, startingPosition)

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
