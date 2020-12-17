package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", CalcBusIDAndWait(input))
	fmt.Println("Part 2: ", CalcEarliestConsecutiveTime(input))
}

func CalcEarliestConsecutiveTime(input []string) int {
	_, buses := parse(input)

	diff := make([]int, 0)
	newBuses := make([]int, 0)
	for i := range buses {
		if buses[i] == 0 {
			continue
		}
		newBuses = append(newBuses, buses[i])
		diff = append(diff, i)
	}
	buses = newBuses

	t := 0
	offset := buses[0]
	cur := 0
	prvT := 0
	newOffset := false
	// count := 0
	newIndex := 0
	for {
		// count++
		t += offset
		found := true
		for i := newIndex; i < len(buses); i++ {
			// count++
			d := t + diff[i]
			r := d % buses[i]
			if r != 0 {
				found = false
				break
			} else {
				if cur == i && newOffset == false {
					fmt.Println("New offset: ", buses[i])
					offset = t - prvT
					newOffset = true

					newIndex = cur + 1
					found = false
					break
				}
				if cur < i {
					newOffset = false
					prvT = t
					offset = buses[i]
					if i < len(buses)-1 {
						found = false
					}
					newIndex = 0
					cur = i
					break
				}
			}
		}
		if found {
			break
		}
	}

	// fmt.Println(count)

	return t
}

func CalcBusIDAndWait(input []string) int {
	depart, buses := parse(input)

	wait := math.MaxInt64
	var busId int
	for _, b := range buses {
		if b == 0 {
			continue
		}

		closeTime := (depart / b) * b
		remain := depart % b

		if remain > 0 {
			closeTime += b
		}

		w := closeTime - depart
		if w < wait {
			wait = w
			busId = b
		}
	}

	return busId * wait
}

func parse(input []string) (int, []int) {
	depart, err := strconv.ParseInt(input[0], 10, 64)
	if err != nil {
		panic("Can't parse earliest depart time")
	}

	s := strings.Split(input[1], ",")
	buses := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == "x" {
			buses[i] = 0
			continue
		}
		b, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			panic("Can't parse bus id")
		}
		buses[i] = int(b)
	}

	return int(depart), buses
}

var input = []string{
	"1000303",
	"41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,541,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,983,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19",
}
