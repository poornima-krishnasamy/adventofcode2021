package main

import (
	"aofc2021/utils"
	_ "embed"
	"fmt"
)

func main() {

	allBinaryCodes := utils.ReadStringInputFile("input-test.txt")

	gamma, epsilon := part1(allBinaryCodes)

	fmt.Println(utils.FromBinStringToInt(gamma) * utils.FromBinStringToInt(epsilon))

	oxygen, co2 := part2(allBinaryCodes)

	fmt.Println(utils.FromBinStringToInt(oxygen) * utils.FromBinStringToInt(co2))

}

func part1(alllines []string) (string, string) {
	gamma, epsilon := "", ""
	for i := 0; i < len(alllines[0]); i++ {
		most, least := findMostLeast(alllines, i)
		//	fmt.Println("most", most, "least", least)
		gamma += string(most)
		epsilon += string(least)
	}
	fmt.Println("gamma", gamma, "epsilon", epsilon)
	return gamma, epsilon
}

func part2(alllines []string) (string, string) {

	oxygen, co2 := "", ""
	oxygen = filterOxygen(alllines, 0, findMostLeast)
	co2 = filterco2(alllines, 0, findMostLeast)
	return oxygen, co2
}

func findMostLeast(lines []string, i int) (uint8, uint8) {
	count_0, count_1 := 0, 0
	for _, line := range lines {
		if line[i] == '0' {
			count_0++
		} else {
			count_1++
		}
	}
	print(" countzero ", count_0, "countone ", count_1, "\n")
	if count_0 > count_1 {
		return '0', '1'
	}
	return '1', '0'
}

func filterOxygen(alllines []string, index uint8, f func([]string, int) (uint8, uint8)) string {
	if len(alllines) <= 1 {
		return alllines[0]
	}
	var newlines []string

	most, _ := f(alllines, int(index))
	for _, s := range alllines {
		if s[index] == most {
			newlines = append(newlines, s)
		}
	}
	return filterOxygen(newlines, index+1, f)
}

func filterco2(alllines []string, index uint8, f func([]string, int) (uint8, uint8)) string {
	if len(alllines) <= 1 {
		return alllines[0]
	}
	var newco2lines []string

	_, least := f(alllines, int(index))

	for _, s := range alllines {
		if s[index] == least {
			newco2lines = append(newco2lines, s)
		}
	}
	if len(newco2lines) > 0 {
		return filterco2(newco2lines, index+1, f)
	} else {
		return filterco2(alllines, index+1, f)
	}
}
