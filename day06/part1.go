package day06

import (
	"errors"
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func printGrid(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

type Orientation int

const (
	North Orientation = iota
	East
	South
	West
)

type Position struct {
	Row         int
	Column      int
	Orientation Orientation
}

func getCurrentPosition(grid [][]string) (Position, error) {
	for i := range grid {
		for j := range grid[i] {
			curr := grid[i][j]
			switch curr {
			case "^":
				return Position{Row: i, Column: j, Orientation: North}, nil
			case ">":
				return Position{Row: i, Column: j, Orientation: East}, nil
			case "v":
				return Position{Row: i, Column: j, Orientation: South}, nil
			case "<":
				return Position{Row: i, Column: j, Orientation: West}, nil

			default:
				continue
			}
		}
	}
	return Position{}, errors.New("No position found in grid")
}

func isObstacle(grid *[][]string, i, j int) bool {
	return (*grid)[i][j] == "#"
}

func turnRight(pos Position) Position {
	return Position{Row: pos.Row, Column: pos.Column, Orientation: (pos.Orientation + 1) % 4}
}

// Returns new position after taking a step or error if we fall out of grid
func step(grid [][]string, currPos Position) (Position, error) {

	targetRow := currPos.Row
	targetCol := currPos.Column

	switch currPos.Orientation {
	case North:
		targetRow = currPos.Row - 1
	case East:
		targetCol = currPos.Column + 1
	case South:
		targetRow = currPos.Row + 1
	case West:
		targetCol = currPos.Column - 1
	}

	// if invalid location -> we have fallen of the grid
	if targetRow < 0 || targetRow >= len(grid) || targetCol < 0 || targetCol >= len(grid) {
		return Position{}, errors.New("Reached end of grid")

	}

	if isObstacle(&grid, targetRow, targetCol) {
		return turnRight(currPos), nil
	}

	// Move ahead to target position
	return Position{
		Row:         targetRow,
		Column:      targetCol,
		Orientation: currPos.Orientation,
	}, nil

}

func markVisisted(grid *[][]string, counter *int, pos Position) {
	(*grid)[pos.Row][pos.Column] = "X"
	(*counter) += 1
}

func wasVisited(grid *[][]string, pos Position) bool {
	return (*grid)[pos.Row][pos.Column] == "X"

}

func Part1() {
	grid, err := utils.Read2dCharArray("day06/input.txt")

	if err != nil {
		fmt.Println("Error reading data as 2d char array")
		return
	}
	locationCounter := 0

	position, err := getCurrentPosition(grid)
	if err != nil {
		return
	}
	// mark starting position as visited
	markVisisted(&grid, &locationCounter, position)

	for true {

		position, err = step(grid, position)
		// fell of the grid
		if err != nil {
			break
		}

		// count unique locations visited
		visited := wasVisited(&grid, position)
		if !visited {
			markVisisted(&grid, &locationCounter, position)
		}

	}

	printGrid(grid)
	fmt.Println("Number of location visited", locationCounter)

}
