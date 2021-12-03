package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dd := readTxtFile("part1.txt")

	fmt.Println(FindPositionPart1(dd))
	fmt.Println(FindPositionPart2(dd))
}

func FindPositionPart1(instructions []string) int {
	var horizontal, depth int

	for _, i := range instructions {
		dir, amnt := ParseInstruction(i)

		switch dir {
		case "forward":
			horizontal += amnt
		case "up":
			depth -= amnt
		case "down":
			depth += amnt
		default:
			panic("unexpected dir " + dir)
		}
	}

	return horizontal * depth
}

func FindPositionPart2(instructions []string) int {
	var horizontal, depth, aim int

	for _, i := range instructions {
		dir, amnt := ParseInstruction(i)

		switch dir {
		case "forward":
			horizontal += amnt
			depth += aim * amnt
		case "up":
			aim -= amnt
		case "down":
			aim += amnt
		default:
			panic("unexpected dir " + dir)
		}
	}

	return horizontal * depth
}

func ParseInstruction(instruction string) (string, int) {
	ss := strings.Split(instruction, " ")
	d, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return ss[0], int(d)
}
