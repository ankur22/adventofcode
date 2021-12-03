package main

import (
	"bufio"
	"os"
)

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
