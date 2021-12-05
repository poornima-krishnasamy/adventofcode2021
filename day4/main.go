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

	parts := strings.SplitN(input, "\n\n", 2)

	drawn := strings.Split(parts[0], ",")

	boards := strings.Split(parts[1], "\n\n")

	fmt.Println("drawn---", drawn)

	part1(drawn, boards)

}

func part1(drawn []string, boards []string) {

	fmt.Println("length=", len(boards))

	//markTheBoard := make([][]int, len(boards))

	drawnInt := make([]int, len(drawn))
	for i := range drawnInt {
		drawnInt[i], _ = strconv.Atoi(drawn[i])
	}

	markedBoard := [100][100]int{}
	for _, call := range drawnInt {
		for i, board := range boards {

			oneBoard := strings.Replace(board, "\n", " ", -1)

			oneBoard = strings.TrimLeft(oneBoard, "")

			fmt.Println("oneBoard", oneBoard)

			oneBoardItems := strings.Split(oneBoard, " ")

			oneBoardInt := make([]int, len(oneBoardItems))
			for i := range oneBoardItems {
				oneBoardInt[i], _ = strconv.Atoi(oneBoardItems[i])
			}

			fmt.Println("oneBoardInt", oneBoardInt)
			markBoard(oneBoardInt, call)

			fmt.Println("onelen", len(oneBoardInt))
			for j, item := range oneBoardInt {
				markedBoard[i][j] = item
			}

			//	fmt.Println(markedBoard[i])

		}
		fmt.Println("next board")
	}

}

func markBoard(board []int, call int) {
	for i, boardItem := range board {
		if boardItem == call {
			board[i] = -1
		}

	}
}
