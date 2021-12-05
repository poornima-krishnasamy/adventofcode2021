package main

import (
	"aofc2021/utils"
	_ "embed"
	"fmt"
	"strings"
)

func main() {

	allPositions := utils.ReadStringInputFile("input.txt")

	fmt.Println("Part One --", part1(allPositions))

	fmt.Println("Part Two --", part2(allPositions))

}

func part1(allPositions []string) int {
	horizontal := 0
	depth := 0

	for _, position := range allPositions {
		step := strings.Fields(position)
		value := utils.ToInt(step[1])
		switch step[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return horizontal * depth
}

func part2(allPositions []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, position := range allPositions {
		step := strings.Fields(position)
		value := utils.ToInt(step[1])
		switch step[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return horizontal * depth
}
