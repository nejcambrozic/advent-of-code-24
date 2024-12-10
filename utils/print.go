package utils

import "fmt"

func Print2dArray(arr [][]string) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}
