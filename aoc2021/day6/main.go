package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := readTxtFile("part1.txt")

	fmt.Print(CountLanternfishPart1(strings.Split(input[0], ","), 80))
	fmt.Print(CountLanternfishPart1(strings.Split(input[0], ","), 256))
}

func CountLanternfishPart1(input []string, daysLeft int) int {
	cyclesDays := 7
	cache := make(map[int]int) // Dynamic Programming! Cache the result :D

	var created int
	for _, initialState := range input {
		i, err := strconv.Atoi(initialState)
		if err != nil {
			panic(err)
		}

		newCreated, ok := cache[i]
		if !ok {
			newCreated = calcCreated(i, daysLeft, cyclesDays)
			cache[i] = newCreated
		} else {
			fmt.Println("cache hit")
		}
		created += newCreated
	}

	return created + len(input)
}

func calcCreated(initialState, initialDaysLeft, cyclesDays int) int {
	additionalDays := cyclesDays - (initialState + 1)
	daysLeft := initialDaysLeft + additionalDays
	created := daysLeft / cyclesDays

	var addCreated int
	for i := 0; i < created; i++ {
		addCreated += calcCreated(8, initialDaysLeft-(initialState+1)-(i*cyclesDays), cyclesDays)
	}

	return created + addCreated
}
