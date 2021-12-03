package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := readTxtFile("part1.txt")

	fmt.Println(GetPowerConsumptionPart1(input))
	fmt.Println(GetCO2O2Part2(input))
}

func GetPowerConsumptionPart1(input []string) int {
	halfSize := len(input) / 2

	sum := make([]int, len(input[0]))

	for _, i := range input {
		for j := 0; j < len(i); j++ {
			v := int(i[j]) - 48 // casting from ascii
			sum[j] += int(v)
		}
	}

	var gamma, epsilon string
	for j := 0; j < len(sum); j++ {
		if sum[j] > halfSize {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(g * e)
}

func GetCO2O2Part2(input []string) int {
	index := 0

	ones := make([]string, 0)
	zeroes := make([]string, 0)
	for _, i := range input {
		if i[index] == 48 { // 48 ascii for 0
			zeroes = append(zeroes, i)
		} else {
			ones = append(ones, i)
		}
	}

	var o2R, co2R int
	if len(zeroes) > len(ones) {
		o2R = getO2Part2(zeroes, index+1)
		co2R = getCo2Part2(ones, index+1)
	} else {
		o2R = getO2Part2(ones, index+1)
		co2R = getCo2Part2(zeroes, index+1)
	}

	return o2R * co2R
}

func getO2Part2(input []string, index int) int {
	if len(input) == 1 {
		v, err := strconv.ParseInt(input[0], 2, 64)
		if err != nil {
			panic(err)
		}
		return int(v)
	}

	ones := make([]string, 0)
	zeroes := make([]string, 0)

	for _, i := range input {
		if i[index] == 48 { // 48 ascii for 0
			zeroes = append(zeroes, i)
		} else {
			ones = append(ones, i)
		}
	}

	var o2R int
	if len(zeroes) > len(ones) {
		o2R = getO2Part2(zeroes, index+1)
	} else {
		o2R = getO2Part2(ones, index+1)
	}

	return o2R
}

func getCo2Part2(input []string, index int) int {
	if len(input) == 1 {
		v, err := strconv.ParseInt(input[0], 2, 64)
		if err != nil {
			panic(err)
		}
		return int(v)
	}

	ones := make([]string, 0)
	zeroes := make([]string, 0)

	for _, i := range input {
		if i[index] == 48 { // 48 ascii for 0
			zeroes = append(zeroes, i)
		} else {
			ones = append(ones, i)
		}
	}

	var o2R int
	if len(zeroes) > len(ones) {
		o2R = getCo2Part2(ones, index+1)
	} else {
		o2R = getCo2Part2(zeroes, index+1)
	}

	return o2R
}
