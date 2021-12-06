package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	state := strings.Split(input, ",")

	part1(state, 80)
	part1(state, 256)

}
func part1(state []string, days int) {

	stateInt := make([]int64, len(state))

	for i := range state {
		stateInt[i], _ = strconv.ParseInt(state[i], 10, 64)
	}

	generateDays(stateInt, days)

}

func generateDays(state []int64, days int) {

	timerState := make([]int64, 9)

	// go vertical to count state of fish on each day
	for _, i := range state {
		timerState[i]++
	}

	for day := 0; day < days; day++ {
		fishToAdd := timerState[0]
		for i := 1; i < len(timerState); i++ {
			timerState[i-1] = timerState[i]
		}
		timerState[6] += fishToAdd
		timerState[8] = fishToAdd
	}

	var total int64 = 0
	for i := 0; i < len(timerState); i++ {
		total += timerState[i]
	}

	fmt.Println(total)

}
