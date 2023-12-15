package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func PuzzleOne() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("error opening puzzle input: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var y int
	for sc.Scan() {
		b := sc.Bytes()
		var byteSlice []byte
		for i, x := range b {
			if x >= '0' && x <= '9' {
				// handle multi digit int check
				byteSlice = append(byteSlice, x)
			}
		}

	}
}
