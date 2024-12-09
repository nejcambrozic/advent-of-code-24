package day05

import (
	"fmt"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func isValidPlan(plan []int, mustBeBefore map[int]mapset.Set[int]) bool {

	visited := mapset.NewSet[int]()
	for _, pageNum := range plan {
		value, ok := mustBeBefore[pageNum]

		// page doesn't have any rules about being before other page
		if !ok {
			visited.Add(pageNum)
			continue
		}

		// plan is not valid if we have already seen any number that
		// should be after current `pageNum`
		inter := value.Intersect(visited)
		if !inter.IsEmpty() {
			return false
		}
		visited.Add(pageNum)
	}
	return true
}

func getInstructionAndUpdatePlans(filePath string) (map[int]mapset.Set[int], [][]int, error) {

	fileContent, err := utils.ReadString("day05/input.txt")
	if err != nil {
		fmt.Println("Error reading file as string", err)
		return nil, nil, err
	}

	// first part above empty line: order instructions
	// second: update page plans
	inputParts := strings.Split(fileContent, "\n\n")
	orderInstructionsInput := inputParts[0]
	updatePlanInput := inputParts[1]

	// Store order instructions as a map.
	// Each key holds a set of numbers where key must occur before all numbers in set
	// So given the instructions
	// ---
	// 1|2
	// 1|3
	// 2|3
	// ---
	// The orderedInstructions will look like: {1: {2,3}, 2: {3}}
	orderInstructions := make(map[int]mapset.Set[int])
	for _, line := range strings.Split(orderInstructionsInput, "\n") {
		orderInstruction := strings.Split(line, "|")
		first, _ := strconv.Atoi(orderInstruction[0])
		second, _ := strconv.Atoi(orderInstruction[1])

		value, ok := orderInstructions[first]
		if !ok {
			orderInstructions[first] = mapset.NewSet(second)
		} else {
			value.Add(second)
		}
	}

	updatePlans, err := utils.Sto2dIntArray(updatePlanInput, ",")

	if err != nil {
		return nil, nil, err
	}

	return orderInstructions, updatePlans, nil
}

func Part1() {

	orderInstructions, updatePlans, err := getInstructionAndUpdatePlans("day05/input.txt")

	if err != nil {
		fmt.Println("Error getting data", err)
		return
	}

	midElementSum := 0
	for _, plan := range updatePlans {
		if isValidPlan(plan, orderInstructions) {
			midEl := plan[len(plan)/2]
			midElementSum += midEl
		}
	}

	fmt.Println("MidElementSum", midElementSum)
}
