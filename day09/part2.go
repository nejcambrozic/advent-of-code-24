package day09

import (
	"errors"
	"fmt"

	"github.com/nejcambrozic/advent-of-code-24-go/utils"
)

func findFileIdx(fileId int, fs []int) (int, int) {

	start := -1
	end := -1
	begin := true
	for i := 0; i < len(fs); i++ {
		if fs[i] == fileId {
			if begin {
				start = i
				begin = false
			}
			end = i
		} else if !begin {
			return start, end
		}
	}

	return start, end
}

func findLastFileIdxs(fs []int) (int, int, int, error) {

	begin := true
	end := -1
	start := -1
	fileId := -1

	for i := len(fs) - 1; i > 0; i-- {
		if begin && fs[i] != -1 {
			fileId = fs[i]
			end = i
			fileId = fs[i]
			begin = false
		}
		if fs[i] == fileId {
			start = i
		} else {
			return start, end, fileId, nil
		}
	}

	return start, end, fileId, errors.New("Not found")
}

func findAvailableEmptySpace(spaceRequired int, fs []int) (int, int, error) {

	begin := true
	start := -1
	end := -1
	counter := 0

	for i := 0; i < len(fs); i++ {
		if fs[i] == -1 {
			if begin {
				start = i
				begin = false
			}
			counter += 1
			end = i

			if counter == spaceRequired {
				return start, end, nil
			}

		} else {
			begin = true
			start = -1
			end = -1
			counter = 0
		}

	}
	return start, end, errors.New("Not found")
}

func Part2() {
	diskMap, _ := utils.ReadString("day09/input.txt")

	fileSystem := diskmapToFS(diskMap)

	fileId := 0
	for j := len(fileSystem) - 1; j > 0; j-- {
		if fileSystem[j] != -1 {
			fileId = fileSystem[j]
			break
		}
	}

	for fileId > 0 {

		fileStart, fileEnd := findFileIdx(fileId, fileSystem)
		fileLen := fileEnd - fileStart + 1

		emptyStart, emptyEnd, err := findAvailableEmptySpace(fileLen, fileSystem)

		if err != nil {
			fileId -= 1
			continue

		}

		if emptyStart < fileStart {

			//fmt.Printf("Moving %d (%d) from [%d,%d] to [%d,%d]\n", fileId, fileLen, fileStart, fileEnd, emptyStart, emptyEnd)
			for i := emptyStart; i <= emptyEnd; i++ {
				fileSystem[i] = fileId
			}
			for i := fileEnd; i >= fileStart; i-- {
				fileSystem[i] = -1
			}
			//PrintFS(fileSystem)
		}

		fileId -= 1
	}

	checksum := Cheksum(fileSystem)
	fmt.Println("Final checksum", checksum)

}

