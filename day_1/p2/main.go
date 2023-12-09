package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("../puzzle_input.txt")
	defer f.Close()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		var foundFrontNum, foundBackNum bool
		for i, j := 0, len(ln)-1; i <= j; {
		}
	}
}
