package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := readTxtFile("part1.txt")

	fmt.Println(FindVentOverlapsPart1(input))
	fmt.Println(FindVentOverlapsPart2(input))
}

func FindVentOverlapsPart1(input []string) int {
	cc := parseInput(input)
	grid := make(map[string]int, 1000)
	var overlapCount int

	for _, c := range cc {
		visitAndMarkOnGrid(&grid, c, c.fromX, c.fromY, &overlapCount)
	}

	return overlapCount
}

func visitAndMarkOnGrid(grid *map[string]int, c coords, curX, curY int, overlapCount *int) {
	a := strconv.Itoa(curY) + ":" + strconv.Itoa(curX)
	(*grid)[a]++

	if (*grid)[a] == 2 {
		(*overlapCount)++
	}

	if curX == c.toX && curY == c.toY {
		return
	}

	switch c.dir {
	case dirNorth:
		curY--
	case dirEast:
		curX++
	case dirSouth:
		curY++
	case dirWest:
		curX--
	default:
		panic("unexpected dir")
	}

	visitAndMarkOnGrid(grid, c, curX, curY, overlapCount)
}

func parseInput(input []string) []coords {
	cc := make([]coords, 0, len(input))
	for _, i := range input {
		i = strings.Replace(i, " -> ", ",", 1)
		si := strings.Split(i, ",")

		if si[0] != si[2] && si[1] != si[3] {
			continue
		}

		c := coords{}

		var err error
		c.fromX, err = strconv.Atoi(si[0])
		if err != nil {
			panic(err)
		}
		c.fromY, err = strconv.Atoi(si[1])
		if err != nil {
			panic(err)
		}
		c.toX, err = strconv.Atoi(si[2])
		if err != nil {
			panic(err)
		}
		c.toY, err = strconv.Atoi(si[3])
		if err != nil {
			panic(err)
		}

		if c.fromX < c.toX {
			c.dir = dirEast
		} else if c.fromX > c.toX {
			c.dir = dirWest
		} else if c.fromY < c.toY {
			c.dir = dirSouth
		} else if c.fromY > c.toY {
			c.dir = dirNorth
		}

		cc = append(cc, c)
	}
	return cc
}

func FindVentOverlapsPart2(input []string) int {
	cc := parseInputPart2(input)
	grid := make(map[string]int, 1000)
	var overlapCount int

	for _, c := range cc {
		visitAndMarkOnGridPart2(&grid, c, c.fromX, c.fromY, &overlapCount)
	}

	return overlapCount
}

func visitAndMarkOnGridPart2(grid *map[string]int, c coords, curX, curY int, overlapCount *int) {
	a := strconv.Itoa(curY) + ":" + strconv.Itoa(curX)
	(*grid)[a]++

	if (*grid)[a] == 2 {
		(*overlapCount)++
	}

	if curX == c.toX && curY == c.toY {
		return
	}

	switch c.dir {
	case dirNorth:
		curY--
	case dirEast:
		curX++
	case dirSouth:
		curY++
	case dirWest:
		curX--
	case dirNorthEast:
		curY--
		curX++
	case dirNorthWest:
		curY--
		curX--
	case dirSouthEast:
		curY++
		curX++
	case dirSouthWest:
		curY++
		curX--
	default:
		panic("unexpected dir")
	}

	visitAndMarkOnGridPart2(grid, c, curX, curY, overlapCount)
}

func parseInputPart2(input []string) []coords {
	cc := make([]coords, 0, len(input))
	for _, i := range input {
		i = strings.Replace(i, " -> ", ",", 1)
		si := strings.Split(i, ",")

		c := coords{}

		var err error
		c.fromX, err = strconv.Atoi(si[0])
		if err != nil {
			panic(err)
		}
		c.fromY, err = strconv.Atoi(si[1])
		if err != nil {
			panic(err)
		}
		c.toX, err = strconv.Atoi(si[2])
		if err != nil {
			panic(err)
		}
		c.toY, err = strconv.Atoi(si[3])
		if err != nil {
			panic(err)
		}

		if c.fromX < c.toX && c.fromY == c.toY {
			c.dir = dirEast
		} else if c.fromX > c.toX && c.fromY == c.toY {
			c.dir = dirWest
		} else if c.fromY < c.toY && c.fromX == c.toX {
			c.dir = dirSouth
		} else if c.fromY > c.toY && c.fromX == c.toX {
			c.dir = dirNorth
		} else if c.fromY > c.toY && c.fromX < c.toX {
			c.dir = dirNorthEast
		} else if c.fromY < c.toY && c.fromX < c.toX {
			c.dir = dirSouthEast
		} else if c.fromY > c.toY && c.fromX > c.toX {
			c.dir = dirNorthWest
		} else if c.fromY < c.toY && c.fromX > c.toX {
			c.dir = dirSouthWest
		}

		cc = append(cc, c)
	}
	return cc
}

type coords struct {
	fromX int
	fromY int
	toX   int
	toY   int
	dir   dirEnum
}

const (
	dirNorth     dirEnum = iota
	dirNorthEast dirEnum = iota
	dirEast      dirEnum = iota
	dirSouthEast dirEnum = iota
	dirSouth     dirEnum = iota
	dirSouthWest dirEnum = iota
	dirWest      dirEnum = iota
	dirNorthWest dirEnum = iota
)

type dirEnum int
