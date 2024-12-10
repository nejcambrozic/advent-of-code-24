package day08

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

type Location struct {
	Row    int
	Column int
}

func isValidLocation(loc Location, gridSize int) bool {
	return loc.Row >= 0 && loc.Row < gridSize && loc.Column >= 0 && loc.Column < gridSize
}

func Part1() {

	grid, _ := utils.Read2dCharArray("day08/input.txt")

	// assuming square grid
	gridSize := len(grid)

	antenas := map[string]mapset.Set[Location]{}
	antizones := mapset.NewSet[Location]()

	for i := range grid {
		for j := range grid {
			freq := grid[i][j]
			if freq != "." {
				loc := Location{Row: i, Column: j}
				locations, ok := antenas[freq]
				if !ok {
					antenas[freq] = mapset.NewSet(loc)
					continue
				}

				// otherwise calculate antizone for each antena for the given frequency already visited
				for existing := range locations.Iter() {

					iDiff := loc.Row - existing.Row
					jDiff := loc.Column - existing.Column

					lookbackPos := Location{Row: existing.Row - iDiff, Column: existing.Column - jDiff}
					if isValidLocation(lookbackPos, gridSize) {
						antizones.Add(lookbackPos)
					}

					lookaheadPos := Location{Row: loc.Row + iDiff, Column: loc.Column + jDiff}
					if isValidLocation(lookaheadPos, gridSize) {
						antizones.Add(lookaheadPos)
					}
				}
				locations.Add(loc)
			}
		}
	}
	fmt.Println("Total unique antizones", antizones.Cardinality())
}
