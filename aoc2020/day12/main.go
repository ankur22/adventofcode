package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", FindManhattanDistance(input))
	fmt.Println("Part 2:", WaypointFindManhattanDistance(input))
}

func WaypointFindManhattanDistance(input []string) int64 {
	r := regexp.MustCompile("(N|S|E|W|L|R|F)([0-9]+)")

	var wayPointX = int64(10)
	var wayPointY = int64(1)

	var x, y int64

	for _, in := range input {
		pin := r.FindStringSubmatch(in)

		if pin == nil {
			panic("Can't parse input")
		}

		dir := pin[1]
		amount, err := strconv.ParseInt(pin[2], 10, 64)
		if err != nil {
			panic("Can't parse int")
		}

		switch dir {
		case "N", "S", "E", "W":
			wayPointX, wayPointY = move(dir, amount, wayPointX, wayPointY)
		case "L":
			rotL := int(amount / 90)
			for r := 0; r < rotL; r++ {
				pX := wayPointX
				pY := wayPointY
				if pX > 0 {
					wayPointY = pX
				} else {
					wayPointY = pX
				}
				if pY > 0 {
					wayPointX = -pY
				} else {
					wayPointX = -pY
				}
			}
		case "R":
			rotL := int(amount / 90)
			for r := 0; r < rotL; r++ {
				pX := wayPointX
				pY := wayPointY
				if pX > 0 {
					wayPointY = -pX
				} else {
					wayPointY = -pX
				}
				if pY > 0 {
					wayPointX = pY
				} else {
					wayPointX = pY
				}
			}
		case "F":
			x += amount * wayPointX
			y += amount * wayPointY
		}
	}

	fmt.Println("x, y: ", x, y)

	return int64(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func FindManhattanDistance(input []string) int64 {
	r := regexp.MustCompile("(N|S|E|W|L|R|F)([0-9]+)")

	var x, y int64
	faces := map[int64]string{0: "N", 90: "E", 180: "S", 270: "W"}
	currentFace := int64(90)

	for _, in := range input {
		pin := r.FindStringSubmatch(in)

		if pin == nil {
			panic("Can't parse input")
		}

		dir := pin[1]
		amount, err := strconv.ParseInt(pin[2], 10, 64)
		if err != nil {
			panic("Can't parse int")
		}

		switch dir {
		case "N", "S", "E", "W":
			x, y = move(dir, amount, x, y)
		case "L":
			cf := float64(currentFace - amount)
			cf /= 360.0
			cf = cf - float64(int64(cf))
			if cf < 0 {
				cf = 1.0 - math.Abs(cf)
			}
			cf *= 360.0
			currentFace = int64(cf)
		case "R":
			cf := float64(currentFace + amount)
			cf /= 360.0
			cf = cf - float64(int64(cf))
			if cf > 1 {
				cf = cf - 1.0
			}
			cf *= 360.0
			currentFace = int64(cf)
		case "F":
			x, y = move(faces[currentFace], amount, x, y)
		}
	}

	fmt.Println("x, y: ", x, y)

	return int64(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func move(dir string, amount, x, y int64) (int64, int64) {
	switch dir {
	case "N":
		y += amount
	case "S":
		y -= amount
	case "E":
		x += amount
	case "W":
		x -= amount
	}
	return x, y
}

var input = []string{
	"W4",
	"N1",
	"F35",
	"R90",
	"F2",
	"N3",
	"W4",
	"S1",
	"R90",
	"F89",
	"W3",
	"N3",
	"R90",
	"W5",
	"F46",
	"L90",
	"W5",
	"L90",
	"F15",
	"W5",
	"N5",
	"E1",
	"F92",
	"N2",
	"W5",
	"L180",
	"F27",
	"N4",
	"F32",
	"R90",
	"S5",
	"F12",
	"E3",
	"R90",
	"N2",
	"F54",
	"R180",
	"E2",
	"L90",
	"E1",
	"R180",
	"F98",
	"N5",
	"F41",
	"S1",
	"R180",
	"N5",
	"F24",
	"W5",
	"R90",
	"S4",
	"R180",
	"W4",
	"R90",
	"W3",
	"N5",
	"E2",
	"F69",
	"N2",
	"R180",
	"F48",
	"W2",
	"N4",
	"E5",
	"S1",
	"W1",
	"F39",
	"S5",
	"L90",
	"E4",
	"R180",
	"E4",
	"L180",
	"E1",
	"R90",
	"F97",
	"S2",
	"L270",
	"N5",
	"E2",
	"R90",
	"S1",
	"R90",
	"E4",
	"N4",
	"F2",
	"N3",
	"L90",
	"E4",
	"F84",
	"W5",
	"N2",
	"E4",
	"S2",
	"E1",
	"S5",
	"E2",
	"N4",
	"E1",
	"F94",
	"R180",
	"N2",
	"L90",
	"W4",
	"L180",
	"F82",
	"E1",
	"L270",
	"F14",
	"S5",
	"S5",
	"R90",
	"F75",
	"R90",
	"E5",
	"S3",
	"R90",
	"N2",
	"F94",
	"L180",
	"E4",
	"F3",
	"E5",
	"N3",
	"R90",
	"N1",
	"E1",
	"S1",
	"R90",
	"N4",
	"L90",
	"E3",
	"L270",
	"F45",
	"E2",
	"S2",
	"E1",
	"F67",
	"L90",
	"F15",
	"N3",
	"W4",
	"N3",
	"R90",
	"F51",
	"N3",
	"S2",
	"R90",
	"N1",
	"L180",
	"W3",
	"L90",
	"W5",
	"F8",
	"N3",
	"F36",
	"N2",
	"L90",
	"S4",
	"F91",
	"N1",
	"E4",
	"R90",
	"E5",
	"L90",
	"E2",
	"F77",
	"R90",
	"E4",
	"F8",
	"W2",
	"R90",
	"N5",
	"R90",
	"E5",
	"S4",
	"W5",
	"F68",
	"N1",
	"F72",
	"S3",
	"W3",
	"F36",
	"E5",
	"F53",
	"E1",
	"E1",
	"N3",
	"W3",
	"N5",
	"L270",
	"S2",
	"L180",
	"E3",
	"S2",
	"W1",
	"R90",
	"E5",
	"S1",
	"W5",
	"N1",
	"F16",
	"W5",
	"N2",
	"E2",
	"N3",
	"F9",
	"S3",
	"L180",
	"F33",
	"E2",
	"L90",
	"R90",
	"F13",
	"S4",
	"W3",
	"S5",
	"F71",
	"N1",
	"S2",
	"F49",
	"E5",
	"F3",
	"S3",
	"L270",
	"N4",
	"E3",
	"R270",
	"N2",
	"L90",
	"F96",
	"S3",
	"R180",
	"F48",
	"E3",
	"N3",
	"F2",
	"E1",
	"S5",
	"F88",
	"W4",
	"F98",
	"E5",
	"F45",
	"R90",
	"N5",
	"W5",
	"L90",
	"S3",
	"E1",
	"F18",
	"E3",
	"S1",
	"F59",
	"L270",
	"N4",
	"L180",
	"W4",
	"S2",
	"F88",
	"L90",
	"N4",
	"F1",
	"L180",
	"E4",
	"L90",
	"W4",
	"R180",
	"F50",
	"L180",
	"W5",
	"F14",
	"S4",
	"F6",
	"R90",
	"E1",
	"N5",
	"E3",
	"F51",
	"R90",
	"E5",
	"R90",
	"F78",
	"S1",
	"W4",
	"N2",
	"R90",
	"F43",
	"E2",
	"N4",
	"R90",
	"F44",
	"L90",
	"F16",
	"N2",
	"W4",
	"S2",
	"F53",
	"S4",
	"W5",
	"F30",
	"L90",
	"E4",
	"F87",
	"E3",
	"L90",
	"S1",
	"F34",
	"W1",
	"L90",
	"F74",
	"S3",
	"F4",
	"E4",
	"N2",
	"F34",
	"W1",
	"L180",
	"F57",
	"E3",
	"R90",
	"W4",
	"N1",
	"R90",
	"E3",
	"R180",
	"F51",
	"S2",
	"R90",
	"S1",
	"E4",
	"S4",
	"F19",
	"W3",
	"F64",
	"E3",
	"N1",
	"E2",
	"L90",
	"E3",
	"W1",
	"S3",
	"R90",
	"F81",
	"E3",
	"S5",
	"R90",
	"F11",
	"R90",
	"W4",
	"N5",
	"W2",
	"S4",
	"W1",
	"S4",
	"F70",
	"L180",
	"F75",
	"F55",
	"S2",
	"W3",
	"S1",
	"L90",
	"E3",
	"L90",
	"W1",
	"F41",
	"S3",
	"L90",
	"N5",
	"L90",
	"W4",
	"F6",
	"R180",
	"E4",
	"F37",
	"L90",
	"F68",
	"R90",
	"W3",
	"F82",
	"S5",
	"L270",
	"S5",
	"F86",
	"L90",
	"F44",
	"R90",
	"W4",
	"R180",
	"S3",
	"F73",
	"L90",
	"E5",
	"R90",
	"E1",
	"S3",
	"L90",
	"F15",
	"R90",
	"F84",
	"L90",
	"R90",
	"S4",
	"W2",
	"E3",
	"R180",
	"F29",
	"R90",
	"W2",
	"F28",
	"N4",
	"E2",
	"N2",
	"F66",
	"R90",
	"E5",
	"R90",
	"L270",
	"S1",
	"F34",
	"E3",
	"S3",
	"F57",
	"N4",
	"S4",
	"F27",
	"N4",
	"W5",
	"L180",
	"E1",
	"S2",
	"F8",
	"S4",
	"E5",
	"L180",
	"F40",
	"L90",
	"N5",
	"W5",
	"N2",
	"L270",
	"F9",
	"E5",
	"N2",
	"E2",
	"L90",
	"N4",
	"R90",
	"S2",
	"E2",
	"F70",
	"E1",
	"N2",
	"F100",
	"L90",
	"E1",
	"F59",
	"L180",
	"E3",
	"S3",
	"L180",
	"F97",
	"W1",
	"F61",
	"S1",
	"E2",
	"N3",
	"R90",
	"W1",
	"E2",
	"R90",
	"F50",
	"S2",
	"W5",
	"F86",
	"S1",
	"L90",
	"S3",
	"F8",
	"L90",
	"E2",
	"F36",
	"L180",
	"E4",
	"R180",
	"E3",
	"N3",
	"E1",
	"F94",
	"N3",
	"F64",
	"L90",
	"S1",
	"S1",
	"F33",
	"E3",
	"S1",
	"L90",
	"N3",
	"W5",
	"N4",
	"L90",
	"F53",
	"W1",
	"N5",
	"F67",
	"E2",
	"F43",
	"N5",
	"W3",
	"W3",
	"L90",
	"L90",
	"E1",
	"S4",
	"W3",
	"R180",
	"S1",
	"F79",
	"S2",
	"W2",
	"F23",
	"L90",
	"W3",
	"F65",
	"E1",
	"R180",
	"E4",
	"N3",
	"L180",
	"S4",
	"E3",
	"E5",
	"L90",
	"F92",
	"W4",
	"F57",
	"S3",
	"L90",
	"S2",
	"L90",
	"S2",
	"F50",
	"E3",
	"L90",
	"S3",
	"R180",
	"E4",
	"F14",
	"E5",
	"S5",
	"E3",
	"N2",
	"L180",
	"W4",
	"S5",
	"S1",
	"L90",
	"F1",
	"S4",
	"F4",
	"R90",
	"S5",
	"F15",
	"N4",
	"L90",
	"W3",
	"L90",
	"S5",
	"L270",
	"W1",
	"R90",
	"S3",
	"L90",
	"W4",
	"F16",
	"R90",
	"S1",
	"F61",
	"L90",
	"N5",
	"R90",
	"F21",
	"R90",
	"N4",
	"W1",
	"F36",
	"S1",
	"L90",
	"W4",
	"R90",
	"E3",
	"N1",
	"F68",
	"N4",
	"E4",
	"R90",
	"E5",
	"S1",
	"R90",
	"F2",
	"R90",
	"E3",
	"W2",
	"R180",
	"F13",
	"R90",
	"S5",
	"F99",
	"E5",
	"R90",
	"F8",
	"S5",
	"F54",
	"L90",
	"S4",
	"R180",
	"W2",
	"S1",
	"F32",
	"W3",
	"N2",
	"W3",
	"N3",
	"E2",
	"N2",
	"R180",
	"S5",
	"E1",
	"R90",
	"F73",
	"L90",
	"W4",
	"R90",
	"S4",
	"F29",
	"F60",
	"L90",
	"N5",
	"E3",
	"N5",
	"E1",
	"F94",
	"N1",
	"E5",
	"R180",
	"N3",
	"E2",
	"R90",
	"W3",
	"L180",
	"F55",
	"N2",
	"W2",
	"R270",
	"F64",
	"E5",
	"W1",
	"N4",
	"E4",
	"N3",
	"E1",
	"R90",
	"E1",
	"S4",
	"L270",
	"W4",
	"S4",
	"L180",
	"F33",
	"R90",
	"E5",
	"S2",
	"W3",
	"W5",
	"N3",
	"W2",
	"F20",
	"W5",
	"S3",
	"W3",
	"R180",
	"F99",
	"E1",
	"F4",
	"E5",
	"F34",
	"F1",
	"F43",
	"S4",
	"E5",
	"R90",
	"N5",
	"F95",
	"W4",
	"F10",
	"N1",
	"L90",
	"S4",
	"W2",
	"L90",
	"F41",
	"L270",
	"F89",
	"L90",
	"F34",
	"N5",
	"F2",
	"W3",
	"F27",
	"S1",
	"R270",
	"W3",
	"N2",
	"R90",
	"F13",
	"N1",
	"F78",
	"W4",
	"R90",
	"S5",
	"F27",
	"N5",
	"E4",
	"F61",
	"L180",
	"N2",
	"R90",
	"W5",
	"S4",
	"S4",
	"E1",
	"F8",
	"R90",
	"E2",
	"R90",
	"E5",
	"F1",
	"E1",
	"L90",
	"F32",
	"W3",
	"S1",
	"S3",
	"R90",
	"E5",
	"F84",
	"N3",
	"F17",
	"R90",
	"S2",
	"F80",
	"L90",
	"F84",
	"S1",
	"F63",
	"R90",
	"N1",
	"R90",
	"N3",
	"F88",
	"F16",
	"N3",
	"W1",
	"N5",
	"L180",
	"S5",
	"F66",
	"R90",
	"W1",
	"F85",
	"S1",
	"F69",
	"R180",
	"S3",
	"F67",
	"F69",
	"W2",
	"F72",
	"R180",
	"W5",
	"F13",
}