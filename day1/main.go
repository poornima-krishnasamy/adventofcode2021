package main

import (
	"aofc2021/utils"
	"fmt"
)

func countIncreased(input []int64) int64 {
	var inc int64

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			inc++
		}
	}

	return inc
}

func main() {
	input := utils.ReadInputFile("input.txt")

	fmt.Println(countIncreased(input))
	var windowedInput []int64
	for i := 2; i < len(input); i++ {
		w := input[i-2] + input[i-1] + input[i]
		windowedInput = append(windowedInput, w)
	}

	fmt.Println(countIncreased(windowedInput))
}
