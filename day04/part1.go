package day04

import (
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func IsMatch(word string) bool {
	return word == "XMAS" || word == "SAMX"
}

func createHitmap(input [][]string) *[][]string {
	rows := len(input)
	cols := len(input[0])
	var hitmap = make([][]string, rows)
	for i := range hitmap {
		hitmap[i] = make([]string, cols)
	}
	for i := range hitmap {
		for j := range hitmap[i] {
			hitmap[i][j] = "."
		}
	}
	return &hitmap
}

func recordHit(row, col [4]int, puzzle *[][]string, hitmap *[][]string) {
	for i := range row {
		(*hitmap)[row[i]][col[i]] = (*puzzle)[row[i]][col[i]]
	}
}

func printHitmap(hitmap *[][]string) {
	fmt.Println("Hitmap")
	hm := *hitmap
	for i := range hm {
		for j := range hm {
			fmt.Print(hm[i][j])
		}
		fmt.Println("")
	}
}

func getWord(iIdxs, jIdxs [4]int, puzzle [][]string) string {
	word := puzzle[iIdxs[0]][jIdxs[0]] + puzzle[iIdxs[1]][jIdxs[1]] + puzzle[iIdxs[2]][jIdxs[2]] + puzzle[iIdxs[3]][jIdxs[3]]
	return word
}

func Part1() {
	puzzle, err := utils.Read2dCharArray("day04/test.txt")
	if err != nil {
		fmt.Println("Error opening file as 2darray")
		return
	}

	hitmap := createHitmap(puzzle)
	foundCount := 0
	for i := range puzzle {
		for j := range puzzle[i] {

			// // horizontal scan
			if j < len(puzzle[i])-3 {
				iIdxs := [4]int{i, i, i, i}
				jIdxs := [4]int{j, j + 1, j + 2, j + 3}

				horizontalWord := getWord(iIdxs, jIdxs, puzzle)
				if IsMatch(horizontalWord) {
					foundCount += 1
					recordHit(iIdxs, jIdxs, &puzzle, hitmap)

				}
			}
			// vertical scan
			if i < len(puzzle)-3 {
				iIdxs := [4]int{i, i + 1, i + 2, i + 3}
				jIdxs := [4]int{j, j, j, j}
				verticalWord := getWord(iIdxs, jIdxs, puzzle)
				if IsMatch(verticalWord) {
					foundCount += 1
					recordHit(iIdxs, jIdxs, &puzzle, hitmap)
				}

			}
			// same diagonal scan
			if i < len(puzzle)-3 && j < len(puzzle[i])-3 {
				iIdxs := [4]int{i, i + 1, i + 2, i + 3}
				jIdxs := [4]int{j, j + 1, j + 2, j + 3}
				rightDiagonalWord := getWord(iIdxs, jIdxs, puzzle)
				if IsMatch(rightDiagonalWord) {
					foundCount += 1
					recordHit(iIdxs, jIdxs, &puzzle, hitmap)
				}
			}
			// opposite diagonal scan
			if i < len(puzzle)-3 && j >= 3 {
				iIdxs := [4]int{i, i + 1, i + 2, i + 3}
				jIdxs := [4]int{j, j - 1, j - 2, j - 3}
				leftDiagonalWord := getWord(iIdxs, jIdxs, puzzle)
				if IsMatch(leftDiagonalWord) {
					foundCount += 1
					recordHit(iIdxs, jIdxs, &puzzle, hitmap)
				}
			}
		}
	}

	printHitmap(hitmap)
	fmt.Println("Total hits count", foundCount)
}
