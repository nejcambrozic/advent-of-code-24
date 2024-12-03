package day2

import (
	"fmt"
	"math"

	utils "github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func isReportSafe(report []int) bool {

	isIncreasing := report[0] < report[len(report)-1]
	for i := 0; i < len(report)-1; i++ {
		first := report[i]
		second := report[i+1]

		diff := math.Abs(float64(first - second))
		diffOk := diff >= 1 && diff <= 3
		if !diffOk {
			return false
		}
		if isIncreasing && (first > second) {
			return false
		} else if !isIncreasing && (second > first) {
			return false
		}
	}
	return true
}

func Part1() {
	reports, err := utils.Read2dArray("day02/test.txt")
	if err != nil {
		fmt.Println("Error reading input as 2dArray", err)
		return
	}

	safeReports := [][]int{}
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports = append(safeReports, report)
		}
	}
	fmt.Println("Num of safe reports", len(safeReports))

}
