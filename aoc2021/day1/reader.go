package main

import (
	"bufio"
	"os"
	"strconv"
)

func readTxtFileInt(filepath string) []int {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	rtn := make([]int, 0)
	bio := bufio.NewScanner(f)
	for bio.Scan() {
		v, err := strconv.ParseInt(bio.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		rtn = append(rtn, int(v))
	}

	return rtn
}

func readTxtFile(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	rtn := make([]string, 0)
	bio := bufio.NewScanner(f)
	for bio.Scan() {
		rtn = append(rtn, bio.Text())
	}

	return rtn
}
