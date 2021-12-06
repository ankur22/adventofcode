package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := readTxtFile("part1.txt")

	fmt.Println(BingoPart1(input))
	fmt.Println(BingoPart2(input))
}

func BingoPart1(input []string) int {
	numbers, boards := parseInput(input)

	for _, n := range numbers {
		for _, b := range boards {
			e, ok := b.numbers[n]
			if !ok {
				continue
			}

			e.picked = true

			if crs := isBingo(b, e); crs != 0 {
				return n * crs
			}
		}
	}

	return 0
}

func BingoPart2(input []string) int {
	numbers, boards := parseInput(input)

	for _, n := range numbers {
		for _, b := range boards {
			e, ok := b.numbers[n]
			if !ok {
				continue
			}

			e.picked = true

			if crs := isBingo(b, e); crs != 0 {
				if isBingoOnAllBoards(boards) {
					return n * crs
				}
			}
		}
	}

	return 0
}

func isBingoOnAllBoards(boards []*Board) bool {
	for _, b := range boards {
		if b.winCount == 0 {
			return false
		}
	}
	return true
}

// Returns 0 if false, else the sum of the row/column
func isBingo(b *Board, e *Element) int {
	pn := getCountNorthIfPicked(b, e)
	ps := getCountSouthIfPicked(b, e)
	// Remove double count
	ps--

	if pn+ps == 5 {
		var sum int
		for _, v := range b.numbers {
			if !v.picked {
				sum += v.value
			}
		}
		b.winCount++
		return sum
	}

	pe := getCountEastIfPicked(b, e)
	pw := getCountWestIfPicked(b, e)
	// Remove double count
	pw--

	if pe+pw == 5 {
		var sum int
		for _, v := range b.numbers {
			if !v.picked {
				sum += v.value
			}
		}
		b.winCount++
		return sum
	}

	return 0
}

func getCountNorthIfPicked(b *Board, e *Element) int {
	if !e.picked {
		return 0
	}

	if e.y == 0 {
		return 1
	}

	p := getCountNorthIfPicked(b, b.grid[e.y-1][e.x])

	return p + 1
}

func getCountSouthIfPicked(b *Board, e *Element) int {
	if !e.picked {
		return 0
	}

	if e.y == 4 {
		return 1
	}

	p := getCountSouthIfPicked(b, b.grid[e.y+1][e.x])

	return p + 1
}

func getCountEastIfPicked(b *Board, e *Element) int {
	if !e.picked {
		return 0
	}

	if e.x == 0 {
		return 1
	}

	p := getCountEastIfPicked(b, b.grid[e.y][e.x-1])

	return p + 1
}

func getCountWestIfPicked(b *Board, e *Element) int {
	if !e.picked {
		return 0
	}

	if e.x == 4 {
		return 1
	}

	p := getCountWestIfPicked(b, b.grid[e.y][e.x+1])

	return p + 1
}

type Element struct {
	x      int
	y      int
	picked bool
	value  int
}

type Board struct {
	numbers  map[int]*Element
	grid     [][]*Element
	winCount int
}

func parseInput(input []string) ([]int, []*Board) {
	calledNumbersS := strings.Split(input[0], ",")
	calledNumbers := make([]int, 0, len(calledNumbersS))
	for _, is := range calledNumbersS {
		i, err := strconv.Atoi(is)
		if err != nil {
			panic(err)
		}

		calledNumbers = append(calledNumbers, i)
	}

	bb := make([]*Board, 0)

	numbers := make(map[int]*Element)
	grid := make([][]*Element, 5)
	y := 0
	for i := 2; i < len(input); i++ {
		line := input[i]
		if line == "" {
			b := Board{
				numbers: numbers,
				grid:    grid,
			}
			bb = append(bb, &b)

			numbers = make(map[int]*Element)
			grid = make([][]*Element, 5)
			y = 0
			continue
		}

		line = strings.Replace(line, "  ", " ", -1)
		lls := strings.Split(line, " ")
		ll := make([]int, 0, len(lls))
		for _, is := range lls {
			if is == "" {
				continue
			}
			i, err := strconv.Atoi(is)
			if err != nil {
				panic(err)
			}

			ll = append(ll, i)
		}

		for x := 0; x < len(ll); x++ {
			e := Element{
				x:     x,
				y:     y,
				value: ll[x],
			}

			numbers[ll[x]] = &e
			grid[y] = append(grid[y], &e)
		}

		y++
	}

	b := Board{
		numbers: numbers,
		grid:    grid,
	}
	bb = append(bb, &b)

	return calledNumbers, bb
}
