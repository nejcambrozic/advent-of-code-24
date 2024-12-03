package day1

import (
	"fmt"
	"math"
	"sort"

	utils "github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func Part1() {

	locationIds, err := utils.Read2dArray("day01/test.txt")
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
	sort.Ints(firstList)
	sort.Ints(secondList)

	distance := 0

	for i := range len(firstList) {
		distance = distance + int(math.Abs(float64(firstList[i])-float64(secondList[i])))

	}
	fmt.Println("Distance", distance)

}
