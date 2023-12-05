package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("/Users/jayreddy/advent_of_code_2023_go/day_1/puzzle_input.txt")
	defer f.Close()
	if err != nil {
		log.Fatalf("error opening puzzle input: %v", err)
	}
	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		var foundI, foundJ bool
		for i, j := 0, len(line)-1; i <= j; {
			if !foundI {
				_, err := strconv.Atoi(string(line[i]))
				if err != nil {
					i++
				} else {
					foundI = true
				}
			}
			if !foundJ {
				_, err := strconv.Atoi(string(line[j]))
				if err != nil {
					j--
				} else {
					foundJ = true
				}
			}
			if foundI && foundJ {
				num, err := strconv.Atoi(string(line[i]) + string(line[j]))
				if err != nil {
					log.Fatalf("error creating addend to sum %v", err)
				}
				sum += num
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	fmt.Printf("total sum is %d", sum)
}
