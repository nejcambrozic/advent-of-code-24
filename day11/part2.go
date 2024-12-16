package day11

import (
	"fmt"
	"strconv"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func cacheKey(stone, blink int) string {
	return fmt.Sprintf("%d-%d", stone, blink)
}

func blinkStone(stone, blinkCount int, cache *map[string]int) int {
	val, ok := (*cache)[cacheKey(stone, blinkCount)]
	if ok {
		return val
	}
	if blinkCount == 0 {
		return 1
	}

	nextBlink := blinkCount - 1
	if stone == 0 {
		stoneCount := blinkStone(1, nextBlink, cache)
		(*cache)[cacheKey(stone, blinkCount)] = stoneCount
		return stoneCount
	}

	numS := strconv.Itoa(stone)
	numLen := len(numS)
	if numLen%2 == 0 {
		left, _ := strconv.Atoi(numS[:numLen/2])
		right, _ := strconv.Atoi(numS[numLen/2:])

		leftCount := blinkStone(left, nextBlink, cache)
		rightCount := blinkStone(right, nextBlink, cache)

		stoneCount := leftCount + rightCount
		(*cache)[cacheKey(stone, blinkCount)] = stoneCount

		return stoneCount
	}

	stoneCount := blinkStone(stone*2024, nextBlink, cache)
	(*cache)[cacheKey(stone, blinkCount)] = stoneCount
	return stoneCount
}

func Part2() {

	stones, _ := utils.ReadIntArray("day11/input.txt")

	cache := map[string]int{}
	total := 0
	blink := 75
	for _, stone := range stones {

		cacheKey := cacheKey(stone, blink)
		val, ok := cache[cacheKey]
		if ok {
			total += val
			continue
		}

		stoneCount := blinkStone(stone, blink, &cache)
		cache[cacheKey] = stoneCount

		total += stoneCount
	}

	fmt.Println("Final number of stones", total)
}
