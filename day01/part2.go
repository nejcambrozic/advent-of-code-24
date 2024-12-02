package day1

import (
	"fmt"

	utils "github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func Part2() {
	locationIds, err := utils.Read2dArray("day01/input.txt")
	if err != nil {
		fmt.Println("Error reading input as 2dArray", err)
		return
	}

	var firstList []int
	var secondList []int

	for _, row := range locationIds {
		firstList = append(firstList, row[0])
		secondList = append(secondList, row[1])
	}

	duplicatedIdsRight := map[int]int{}

	for _, locId := range secondList {
		duplicatedIdsRight[locId] += 1
	}

	similarityScore := 0

	for _, locId := range firstList {
		similarityScore += locId * duplicatedIdsRight[locId]
	}

	fmt.Println("Similarity score", similarityScore)

}
