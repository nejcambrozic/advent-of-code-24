package day08

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func Part2() {

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
				antizones.Add(loc)

				locations, ok := antenas[freq]
				if !ok {
					antenas[freq] = mapset.NewSet(loc)
					continue
				}

				// otherwise calculate antizone for each antena for the given frequency already visited
				for existing := range locations.Iter() {
					iDiff := loc.Row - existing.Row
					jDiff := loc.Column - existing.Column

					for z := 1; z < gridSize; z++ {
						lookbackPos := Location{Row: existing.Row - z*iDiff, Column: existing.Column - z*jDiff}
						if IsValidLocation(lookbackPos, gridSize) {
							antizones.Add(lookbackPos)
						} else {
							break
						}
					}

					for z := 1; z < gridSize; z++ {
						lookaheadPos := Location{Row: loc.Row + z*iDiff, Column: loc.Column + z*jDiff}
						if IsValidLocation(lookaheadPos, gridSize) {
							antizones.Add(lookaheadPos)
						} else {
							break
						}
					}

				}
				locations.Add(loc)
			}
		}
	}
	fmt.Println("Total unique antizones", antizones.Cardinality())
}
