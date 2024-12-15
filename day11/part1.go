package day11

import (
	"fmt"
	"strconv"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func mutateStone(stone int) []int {

	if stone == 0 {
		return []int{1}
	}
	numS := strconv.Itoa(stone)
	numLen := len(numS)
	if numLen%2 == 0 {
		left, _ := strconv.Atoi(numS[:numLen/2])
		right, _ := strconv.Atoi(numS[numLen/2:])
		return []int{left, right}
	}
	return []int{stone * 2024}

}

func Part1() {

	stones, _ := utils.ReadIntArray("day11/test.txt")
	fmt.Println(stones)

	maxBlink := 25

	for blink := range maxBlink {
		newStones := []int{}

		for _, stone := range stones {
			mutated := mutateStone(stone)

			newStones = append(newStones, mutated...)
		}

		fmt.Printf("After %d blink\n", blink)
		//fmt.Println(newStones)
		stones = newStones
	}

	fmt.Println("Final number of stones", len(stones))
}
