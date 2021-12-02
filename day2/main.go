package main

import (
	"aofc2021/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var allPositions string

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

		switch step[0] {
		case "forward":
			horizontal += toInt(step[1])
		case "down":
			depth += toInt(step[1])
		case "up":
			depth -= toInt(step[1])
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
		value := toInt(step[1])
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

func toInt(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return value
}
