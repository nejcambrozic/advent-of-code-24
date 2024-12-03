package day03

import (
	"fmt"
	"regexp"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func Part2() {
	instructions, err := utils.ReadString("day03/test.txt")
	if err != nil {
		fmt.Println("Error reading input as single string", err)
		return
	}

	// We are only looking for instructions between [enabled, disabled] chars
	// in order to handle beginning and end of string let's add enable at the start
	// and disable char at the end
	instructions = "do()" + instructions + "don't()"

	reEnables := regexp.MustCompile(`do\(\)`)
	enabledIdxs := reEnables.FindAllStringIndex(instructions, -1)

	reDisables := regexp.MustCompile(`don't\(\)`)
	disabledIdxs := reDisables.FindAllStringIndex(instructions, -1)

	// pointers pointing towards indexes we are using
	enablePointer := 0
	disablePointer := 0

	totalSum := 0
	disabledUntil := 0
	for enablePointer < len(enabledIdxs) && disablePointer < len(disabledIdxs) {
		startIdx := enabledIdxs[enablePointer][0]
		endIdx := disabledIdxs[disablePointer][0]

		// We found multiple disable instructions in a disabled block,
		// let's forward until the next `disable` instructions which occurs AFTER the current `enable` instruction
		if startIdx > endIdx {
			disablePointer += 1
			continue
		}
		// We found multiple enable instructions in a enabled block,
		// lets forward until next `enable` instruction which occurs AFTER latest `disable` instructions we used.
		if startIdx < disabledUntil {
			enablePointer += 1
			continue
		}

		// Valid enabled instructions
		enabledInstructions := instructions[startIdx:endIdx]
		totalSum += MulSum(enabledInstructions)

		// We used up this pair of indexes -> move to next available
		enablePointer += 1
		disablePointer += 1
		// Store the last disabled idx so we don't over-count multiple `enable` instructions
		// This handles cases like: 'do()1..do()2..don't` -> we must only use do()1..don't once and not also do()2..don't
		// Se the `if` statment above using `disabledUntil`
		disabledUntil = endIdx

	}
	fmt.Println("Enabled sum", totalSum)
}
