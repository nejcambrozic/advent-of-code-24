package day2

import (
	"fmt"
	"math"

	utils "github.com/nejcambrozic/advent-of-code-24-go/utils"
)

// Removes element at index `idx` from `report` without modfying `report` in place
func safeElementRemoval(report []int, idx int) []int {
	reportNoFirstEl := make([]int, len(report))
	copy(reportNoFirstEl, report)

	return append(reportNoFirstEl[:idx], reportNoFirstEl[idx+1:]...)
}

// Validates if report is safe where we allow `levelsAllowedToSkip` levels to be skipped
func isReportSafeAllowSkip(report []int, levelsAllowedToSkip int) bool {

	isIncreasing := report[0] < report[len(report)-1]

	isValid := true
	for i := 0; i < len(report)-1; i++ {
		first := report[i]
		second := report[i+1]

		diff := math.Abs(float64(first - second))
		diffOk := diff >= 1 && diff <= 3
		if !diffOk {
			// fmt.Println("Diff unsafe ", first, second)
			isValid = false
		} else if isIncreasing && (first > second) {
			// fmt.Println("Not increasing ", first, second)
			isValid = false
		} else if !isIncreasing && (second > first) {
			// fmt.Println("Not decreasing ", first, second)
			isValid = false
		}

		if !isValid {
			if levelsAllowedToSkip == 0 {
				return false
			} else {
				decreasedLevel := levelsAllowedToSkip - 1

				removeFirstReport := safeElementRemoval(report, i)
				isFirstValid := isReportSafeAllowSkip(removeFirstReport, decreasedLevel)

				removeSecondReport := safeElementRemoval(report, i+1)
				isSecondValid := isReportSafeAllowSkip(removeSecondReport, decreasedLevel)
				return isFirstValid || isSecondValid
			}
		}
	}
	return true
}

func Part2() {
	reports, err := utils.Read2dArray("day02/test.txt")
	if err != nil {
		fmt.Println("Error reading input as 2dArray", err)
		return
	}

	safeReports := [][]int{}
	for _, report := range reports {
		if isReportSafeAllowSkip(report, 1) {
			safeReports = append(safeReports, report)
		}
	}
	fmt.Println("Num of safe reports", len(safeReports))

}
