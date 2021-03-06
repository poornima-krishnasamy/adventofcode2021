package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadIntInputFile(file string) []int64 {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input []int64
	for scanner.Scan() {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		input = append(input, value)
	}

	return input
}

func ReadStringInputFile(file string) []string {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var line []string
	for scanner.Scan() {
		line = append(line, strings.TrimRight(scanner.Text(), "\n"))
	}
	return line
}

func ToInt(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return value
}

func FromBinStringToInt(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
