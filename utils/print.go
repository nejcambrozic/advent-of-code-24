package utils

import "fmt"

func Print2dArray(arr interface{}) {

	switch v := arr.(type) {

	case [][]int:
		for i := range v {
			for j := range v[i] {
				fmt.Print(v[i][j])
			}
			fmt.Println()
		}
	case [][]string:
		for i := range v {
			for j := range v[i] {
				fmt.Print(v[i][j])
			}
			fmt.Println()
		}
	}
}
