package day05

import (
	"fmt"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
)

func orderPlan(plan []int, orderInstructions map[int]mapset.Set[int]) []int {
	visited := mapset.NewSet[int]()
	for i, pageNum := range plan {
		value, ok := orderInstructions[pageNum]

		// page doesn't have any rules about being before other page
		if !ok {
			visited.Add(pageNum)
			continue
		}

		// if intersection is not empty, we need to switch the order of instructions
		inter := value.Intersect(visited)
		if !inter.IsEmpty() {
			sValue, _ := inter.Pop()
			second := slices.Index(plan, sValue)

			plan[i], plan[second] = plan[second], plan[i]
			return orderPlan(plan, orderInstructions)
		}
		visited.Add(pageNum)
	}
	return plan
}

func Part2() {

	orderInstructions, updatePlans, err := getInstructionAndUpdatePlans("day05/test.txt")

	if err != nil {
		fmt.Println("Error getting data", err)
		return
	}

	midElementSum := 0
	for _, plan := range updatePlans {
		if !isValidPlan(plan, orderInstructions) {

			orderedPlan := orderPlan(plan, orderInstructions)

			midEl := orderedPlan[len(orderedPlan)/2]
			midElementSum += midEl
		}
	}

	fmt.Println("MidElementSum", midElementSum)
}
