package main

import (
	"fmt"
)

func main() {
	dd := readTxtFileInt("part1.txt")

	fmt.Println(FindDepthChangesPart1(dd))
	fmt.Println(FindDepthChangesPart2(dd))
}

func FindDepthChangesPart1(depths []int) int {
	last := 100000000000000
	var count int

	for _, d := range depths {
		if d > last {
			count++
		}
		last = d
	}

	return count
}

func FindDepthChangesPart2(depths []int) int {
	var increases int

	for i := 0; i < len(depths)-3; i++ {
		shared := depths[i+1] + depths[i+2]
		first := depths[i] + shared
		second := shared + depths[i+3]

		if second > first {
			increases++
		}
	}

	return increases
}
