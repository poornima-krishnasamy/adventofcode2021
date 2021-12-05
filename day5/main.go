package main

import (
	"aofc2021/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Coords struct {
	x, y int
}

func main() {

	coords := strings.Replace(input, " -> ", ",", -1)

	coordsSet := strings.Split(coords, "\n")

	line := part1(coordsSet)

	var danger int
	for _, v := range line {
		if v > 1 {
			danger++
		}
	}

	fmt.Println("Part1", danger)

	diagonal := part2(coordsSet)

	var overlap int
	for _, v := range diagonal {
		if v > 1 {
			overlap++
		}
	}
	fmt.Println("Part1", overlap)

}
func part1(coordsSet []string) map[Coords]int {
	line := make(map[Coords]int)

	for _, c := range coordsSet {
		pos := strings.Split(c, ",")
		p1 := Coords{utils.ToInt(pos[0]), utils.ToInt(pos[1])}
		p2 := Coords{utils.ToInt(pos[2]), utils.ToInt(pos[3])}

		nx := p1.x
		ny := p1.y
		if (nx == p2.x) || (ny == p2.y) {
			for {
				pos := Coords{nx, ny}
				line[pos] += 1
				if nx == p2.x && ny == p2.y {
					break
				}
				nx = nextPos(nx, p2.x)
				ny = nextPos(ny, p2.y)
				//	fmt.Println(line[pos])
			}
		}
	}

	return line
}

func part2(coordsSet []string) map[Coords]int {
	diagonal := make(map[Coords]int)

	for _, c := range coordsSet {
		pos := strings.Split(c, ",")
		p1 := Coords{utils.ToInt(pos[0]), utils.ToInt(pos[1])}
		p2 := Coords{utils.ToInt(pos[2]), utils.ToInt(pos[3])}

		nx := p1.x
		ny := p1.y
		for {
			pos := Coords{nx, ny}
			diagonal[pos] += 1
			if nx == p2.x && ny == p2.y {
				break
			}
			nx = nextPos(nx, p2.x)
			ny = nextPos(ny, p2.y)
			//	fmt.Println(diagonal[pos])
		}
	}

	return diagonal
}

func nextPos(n int, last int) int {
	if n < last {
		return n + 1
	} else if n > last {
		return n - 1
	} else {
		return n
	}
}
