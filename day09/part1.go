package day09

import (
	"fmt"
	"strconv"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func diskmapToFS(diskmap string) []int {
	result := []int{}

	fileId := 0
	id := 0
	for i, char := range diskmap {

		num, err := strconv.Atoi(string(char))
		if err != nil {
			return result
		}
		// file length
		if i%2 == 0 {
			id = fileId
			fileId += 1
		} else {
			//space length
			id = -1
		}

		for j := 0; j < num; j++ {
			result = append(result, id)
		}
	}
	return result
}

func PrintFS(fileSys []int) {
	for _, id := range fileSys {
		if id == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(id)
		}
	}
	fmt.Println()
}

func Cheksum(fileSys []int) int {

	sum := 0
	for i, v := range fileSys {
		if v == -1 {
			continue
		}
		sum += i * v
	}
	return sum
}

func Part1() {
	diskMap, _ := utils.ReadString("day09/input.txt")

	fileSystem := diskmapToFS(diskMap)
	//fmt.Println("fileSystem: ")
	//PrintFS(fileSystem)

	emptySpacePointer := 0
	lastFilePointer := len(fileSystem) - 1

	for emptySpacePointer < lastFilePointer {

		if fileSystem[emptySpacePointer] != -1 {
			emptySpacePointer += 1
			continue
		}
		if fileSystem[lastFilePointer] == -1 {
			lastFilePointer -= 1
			continue
		}

		fileSystem[emptySpacePointer] = fileSystem[lastFilePointer]
		fileSystem[lastFilePointer] = -1

	}

	checksum := Cheksum(fileSystem)
	fmt.Println("Final checksum", checksum)

}

