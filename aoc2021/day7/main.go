package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := readTxtFile("part1.txt")
	iis := strings.Split(input[0], ",")

	ii := make([]int, 0, len(iis))
	for _, is := range iis {
		i, err := strconv.Atoi(is)
		if err != nil {
			panic(err)
		}

		ii = append(ii, i)
	}

	fmt.Println(FindHorizLeastCost(ii))
	fmt.Println(FindHorizLeastCostPart2(ii))
}

func FindHorizLeastCost(input []int) int {
	minFuel := 1000000000
	minHori := -1
	prvFuel := 1000000000
	for j := 0; j < 100000; j++ {
		var fuel int
		for _, i := range input {
			fuel += int(math.Abs(float64(i - j)))
		}
		if fuel > prvFuel {
			break
		}
		if fuel < minFuel {
			minFuel = fuel
			minHori = j
		}
		prvFuel = fuel
	}

	fmt.Println(minHori)
	fmt.Println(minFuel)

	return minFuel
}

func FindHorizLeastCostPart2(input []int) int {
	minFuel := 1000000000
	minHori := -1
	prvFuel := 1000000000
	for j := 0; j < 100000; j++ {
		var fuel int
		for _, i := range input {
			diff := int(math.Abs(float64(i - j)))
			actualDiff := 0
			it := 1
			for j := 0; j < diff; j++ {
				actualDiff += it
				it++
			}
			fuel += actualDiff
		}
		if fuel > prvFuel {
			break
		}
		if fuel < minFuel {
			minFuel = fuel
			minHori = j
		}
		prvFuel = fuel
	}

	fmt.Println(minHori)
	fmt.Println(minFuel)

	return minFuel
}
