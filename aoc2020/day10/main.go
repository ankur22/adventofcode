package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Part 1: ", FindJoltageDifference(input))
	fmt.Println("Part 2: ", FindAllCombinations(input))
}

func FindAllCombinations(input []int) int {
	input = sortAndAdd(input)

	m := make(map[int]int, len(input))
	for _, v := range input {
		m[v] = 0
	}

	m[0] = 1
	for _, cur := range input {
		po := cur + 1
		pt := cur + 2
		ptr := cur + 3

		if _, ok := m[po]; ok {
			m[po] += m[cur]
		}

		if _, ok := m[pt]; ok {
			m[pt] += m[cur]
		}

		if _, ok := m[ptr]; ok {
			m[ptr] += m[cur]
		}
	}

	return m[input[len(input)-1]]
}

func FindJoltageDifference(input []int) int {
	input = sortAndAdd(input)

	var dif1, dif3 int
	for i := 1; i < len(input); i++ {
		dif := input[i] - input[i-1]
		switch dif {
		case 1:
			dif1++
		case 3:
			dif3++
		}
	}

	return dif1 * dif3
}

func sortAndAdd(input []int) []int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	input = append(make([]int, 1), input...)
	input = append(input, make([]int, 1)...)

	const chargingOutlet = 0
	const deviceAdapterPlusJoltage = 3

	input[0] = chargingOutlet
	input[len(input)-1] = input[len(input)-2] + deviceAdapterPlusJoltage

	return input
}

var input = []int{
	144,
	10,
	75,
	3,
	36,
	80,
	143,
	59,
	111,
	133,
	1,
	112,
	23,
	62,
	101,
	137,
	41,
	24,
	8,
	121,
	35,
	105,
	161,
	69,
	52,
	21,
	55,
	29,
	135,
	142,
	38,
	108,
	141,
	115,
	68,
	7,
	98,
	82,
	9,
	72,
	118,
	27,
	153,
	140,
	61,
	90,
	158,
	102,
	28,
	134,
	91,
	2,
	17,
	81,
	31,
	15,
	120,
	20,
	34,
	56,
	4,
	44,
	74,
	14,
	147,
	11,
	49,
	128,
	16,
	99,
	66,
	47,
	125,
	155,
	130,
	37,
	67,
	54,
	60,
	48,
	136,
	89,
	119,
	154,
	122,
	129,
	163,
	73,
	100,
	85,
	95,
	30,
	76,
	162,
	22,
	79,
	88,
	150,
	53,
	63,
	92,
}
